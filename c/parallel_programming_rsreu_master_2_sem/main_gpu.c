#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <CL/cl.h>
# define CLOCK_REALTIME			0

#include "utils/bitmap.h"
#include "utils/kernel.h"


#define CHECK_CL(err, msg) \
    if (err != CL_SUCCESS) { \
        fprintf(stderr, "%s: error %d\n", msg, err); \
        exit(1); \
    }

const char* cl_kernel_source =
"__kernel void sobel_rms_kernel(\n"
"    __global const uchar* input,\n"
"    __global uchar* output,\n"
"    __constant float* kern1,\n"
"    __constant float* kern2,\n"
"    int width,\n"
"    int height,\n"
"    int kern_size)\n"
"{\n"
"    int col = get_global_id(0);\n"
"    int row = get_global_id(1);\n"
"\n"
"    if (col >= width || row >= height) return;\n"
"\n"
"    float sum1 = 0.0f;\n"
"    float sum2 = 0.0f;\n"
"    int k_half = kern_size / 2;\n"
"\n"
"    for (int ky = -k_half; ky <= k_half; ky++) {\n"
"        for (int kx = -k_half; kx <= k_half; kx++) {\n"
"            int px = col + kx;\n"
"            int py = row + ky;\n"
"\n"
"            if (px < 0 || px >= width || py < 0 || py >= height) continue;\n"
"\n"
"            int idx = (py * width + px) * 3;\n"
"            uchar b = input[idx];\n"
"            uchar g = input[idx + 1];\n"
"            uchar r = input[idx + 2];\n"
"\n"
"            float pixel_val = 0.299f * r + 0.587f * g + 0.114f * b;\n"
"\n"
"            int k_idx = (ky + k_half) * kern_size + (kx + k_half);\n"
"            sum1 += pixel_val * kern1[k_idx];\n"
"            sum2 += pixel_val * kern2[k_idx];\n"
"        }\n"
"    }\n"
"\n"
"    float rms = sqrt(sum1 * sum1 + sum2 * sum2);\n"
"    rms = clamp(rms, 0.0f, 255.0f);\n"
"    uchar val = (uchar)rms;\n"
"\n"
"    int out_idx = (row * width + col) * 3;\n"
"    output[out_idx]     = val;\n"
"    output[out_idx + 1] = val;\n"
"    output[out_idx + 2] = val;\n"
"}\n";

int main(int argc, char *argv[])
{
    if (argc < 4) {
        fprintf(stderr, "Usage: %s [INPUT_FILE] [OUTPUT_FILE] [ITERATIONS_COUNT]\n", argv[0]);
        return 1;
    }

    FILE* inputImageFile = fopen(argv[1], "rb");
    FILE* outputImageFile = fopen(argv[2], "wb");
    int64_t iterations_count = atoi(argv[3]);

    if (iterations_count <= 0) {
        fprintf(stderr, "Iterations count must be > 0!\n");
        return 1;
    }
    if (!inputImageFile || !outputImageFile) {
        fprintf(stderr, "Can't open image files!\n");
        return 4;
    }

    // ====================== Чтение BMP ======================
    struct BitmapFileHeader bitmapInputFileHeader;
    struct BitmapInfoHeaderV3 bitmapInputInfoHeaderV3;

    fread(&bitmapInputFileHeader, sizeof(struct BitmapFileHeader), 1, inputImageFile);
    fread(&bitmapInputInfoHeaderV3, sizeof(struct BitmapInfoHeaderV3), 1, inputImageFile);
    fseek(inputImageFile, bitmapInputFileHeader.bfOffBits, SEEK_SET);

    if (bitmapInputFileHeader.bfType != 0x4d42 || bitmapInputInfoHeaderV3.biBitCount != 24) {
        fprintf(stderr, "Only 24-bit BMP files supported!\n");
        return 5;
    }

    const int32_t width  = bitmapInputInfoHeaderV3.biWidth;
    const int32_t height = bitmapInputInfoHeaderV3.biHeight;

    const size_t pixels_count = (size_t)width * height;
    struct Bitmap24Pixel* input_image_pixels  = malloc(pixels_count * sizeof(struct Bitmap24Pixel));
    struct Bitmap24Pixel* output_image_pixels = malloc(pixels_count * sizeof(struct Bitmap24Pixel));

    const size_t row_bytes_without_padding = width * sizeof(struct Bitmap24Pixel);
    const size_t padding_bytes = get_row_size_with_padding(row_bytes_without_padding) - row_bytes_without_padding;

    for (size_t row = 0; row < height; row++) {
        fread(input_image_pixels + row * width, sizeof(struct Bitmap24Pixel), width, inputImageFile);
        fseek(inputImageFile, padding_bytes, SEEK_CUR);
    }
    fclose(inputImageFile);

    // ====================== OpenCL инициализация ======================
    cl_platform_id platform;
    cl_device_id device;
    cl_context context;
    cl_command_queue queue;
    cl_program program;
    cl_kernel kernel;
    cl_int err;

    clGetPlatformIDs(1, &platform, NULL);
    err = clGetDeviceIDs(platform, CL_DEVICE_TYPE_GPU, 1, &device, NULL);
    if (err != CL_SUCCESS)
        clGetDeviceIDs(platform, CL_DEVICE_TYPE_DEFAULT, 1, &device, NULL);

    context = clCreateContext(NULL, 1, &device, NULL, NULL, &err); CHECK_CL(err, "clCreateContext");
    queue   = clCreateCommandQueue(context, device, 0, &err); CHECK_CL(err, "clCreateCommandQueue");

    program = clCreateProgramWithSource(context, 1, &cl_kernel_source, NULL, &err);
    CHECK_CL(err, "clCreateProgramWithSource");

    err = clBuildProgram(program, 1, &device, NULL, NULL, NULL);
    if (err != CL_SUCCESS) {
        char log[4096];
        clGetProgramBuildInfo(program, device, CL_PROGRAM_BUILD_LOG, sizeof(log), log, NULL);
        fprintf(stderr, "Build log:\n%s\n", log);
        return 1;
    }

    kernel = clCreateKernel(program, "sobel_rms_kernel", &err); CHECK_CL(err, "clCreateKernel");

    // ====================== Буферы ======================
    struct Kernel k_first = get_sobel_horizontal_kernel();
    struct Kernel k_second = get_sobel_vertical_kernel();

    cl_mem d_input   = clCreateBuffer(context, CL_MEM_READ_ONLY  | CL_MEM_COPY_HOST_PTR,
                                      pixels_count * 3, input_image_pixels, &err);
    cl_mem d_output  = clCreateBuffer(context, CL_MEM_WRITE_ONLY, pixels_count * 3, NULL, &err);
    cl_mem d_kern1   = clCreateBuffer(context, CL_MEM_READ_ONLY  | CL_MEM_COPY_HOST_PTR,
                                      k_first.width * k_first.width * sizeof(float), k_first.data, &err);
    cl_mem d_kern2   = clCreateBuffer(context, CL_MEM_READ_ONLY  | CL_MEM_COPY_HOST_PTR,
                                      k_second.width * k_second.width * sizeof(float), k_second.data, &err);

    CHECK_CL(err, "Buffer creation");

    // ====================== Вычисления ======================
    printf("Started OpenCL calculations (original Sobel logic): iterations = %" PRId64 "\n", iterations_count);

    struct timespec start, end;
    clock_gettime(CLOCK_REALTIME, &start);

    size_t global_work_size[2] = { (size_t)width, (size_t)height };
    size_t local_work_size[2] = { 32, 4 };

    for (int64_t iter = 0; iter < iterations_count; iter++) {
        clSetKernelArg(kernel, 0, sizeof(cl_mem), &d_input);
        clSetKernelArg(kernel, 1, sizeof(cl_mem), &d_output);
        clSetKernelArg(kernel, 2, sizeof(cl_mem), &d_kern1);
        clSetKernelArg(kernel, 3, sizeof(cl_mem), &d_kern2);
        clSetKernelArg(kernel, 4, sizeof(int),    &width);
        clSetKernelArg(kernel, 5, sizeof(int),    &height);
        clSetKernelArg(kernel, 6, sizeof(int),    &k_first.width);

        clEnqueueNDRangeKernel(queue, kernel, 2, NULL, global_work_size, 0, 0, NULL, NULL);
    }

    clFinish(queue);

    clock_gettime(CLOCK_REALTIME, &end);
    double elapsed_s = (end.tv_sec - start.tv_sec);
    elapsed_s += (end.tv_nsec - start.tv_nsec) / 1'000'000'000.0;
    printf("Done!\n");
    printf("Elapsed time: %.3lf s for all iterations;\n", elapsed_s);
    printf("              %.3lf s for one iterations.\n", elapsed_s / iterations_count);

    // ====================== Чтение результата ======================
    clEnqueueReadBuffer(queue, d_output, CL_TRUE, 0,
                        pixels_count * 3, output_image_pixels, 0, NULL, NULL);

    // ====================== Запись BMP ======================
    struct Bitmap24Image bitmap24_output_image = get_initialized_bitmap24_image(width, height, output_image_pixels);

    fwrite(&bitmap24_output_image.bitmap_file_header, sizeof(struct BitmapFileHeader), 1, outputImageFile);
    fwrite(&bitmap24_output_image.bitmap_info_header_v3, sizeof(struct BitmapInfoHeaderV3), 1, outputImageFile);

    char padding_junk[4] = {0};
    for (size_t row = 0; row < height; row++) {
        fwrite(output_image_pixels + row * width, sizeof(struct Bitmap24Pixel), width, outputImageFile);
        fwrite(padding_junk, 1, padding_bytes, outputImageFile);
    }

    fclose(outputImageFile);

    // ====================== Освобождение ======================
    clReleaseMemObject(d_input);
    clReleaseMemObject(d_output);
    clReleaseMemObject(d_kern1);
    clReleaseMemObject(d_kern2);
    clReleaseKernel(kernel);
    clReleaseProgram(program);
    clReleaseCommandQueue(queue);
    clReleaseContext(context);

    free(input_image_pixels);
    free(output_image_pixels);

    return 0;
}


size_t get_row_size_with_padding(const size_t rowSizeWithoutPadding) {
    return ((rowSizeWithoutPadding + 3) &~ 3);
}

int64_t divide_with_ceil(const int64_t a, const int64_t b) {
    return (a + b - 1) / b;
}

struct Bitmap24Image get_initialized_bitmap24_image(const size_t width, const size_t height, struct Bitmap24Pixel* pixels) {
    struct Bitmap24Image bitmap24Image;

    bitmap24Image.width = width;
    bitmap24Image.height = height;

    const size_t rowSizeWithoutPadding = width * sizeof(struct Bitmap24Pixel);
    const size_t rowSizeWithPadding = get_row_size_with_padding(rowSizeWithoutPadding);
    const size_t imageSize = rowSizeWithPadding * height;

    bitmap24Image.bitmap_file_header.bfType = 0x4d42;
    bitmap24Image.bitmap_file_header.bfSize = sizeof(struct BitmapFileHeader) + sizeof(struct BitmapInfoHeaderV3) + imageSize;
    bitmap24Image.bitmap_file_header.bfReserved1 = 0;
    bitmap24Image.bitmap_file_header.bfReserved2 = 0;
    bitmap24Image.bitmap_file_header.bfOffBits = sizeof(struct BitmapFileHeader) + sizeof(struct BitmapInfoHeaderV3);

    bitmap24Image.bitmap_info_header_v3.biWidth = (int32_t)width;
    bitmap24Image.bitmap_info_header_v3.biHeight = (int32_t)height;
    bitmap24Image.bitmap_info_header_v3.biSize = sizeof(struct BitmapInfoHeaderV3);
    bitmap24Image.bitmap_info_header_v3.biPlanes = 1;
    bitmap24Image.bitmap_info_header_v3.biBitCount = 24;
    bitmap24Image.bitmap_info_header_v3.biCompression = 0;
    bitmap24Image.bitmap_info_header_v3.biSizeImage = imageSize;
    bitmap24Image.bitmap_info_header_v3.biXPelsPerMeter = 2835;
    bitmap24Image.bitmap_info_header_v3.biYPelsPerMeter = 2835;
    bitmap24Image.bitmap_info_header_v3.biClrUsed = 0;
    bitmap24Image.bitmap_info_header_v3.biClrImportant = 0;

    bitmap24Image.pixels = pixels;

    return bitmap24Image;
}


struct Kernel get_sobel_vertical_kernel() {
    struct Kernel kernel = {.width = 3, .height = 3};
    kernel.data[0][0] = -1;
    kernel.data[0][1] = -2;
    kernel.data[0][2] = -1;
    kernel.data[1][0] = 0;
    kernel.data[1][1] = 0;
    kernel.data[1][2] = 0;
    kernel.data[2][0] = 1;
    kernel.data[2][1] = 2;
    kernel.data[2][2] = 1;
    return kernel;
}

struct Kernel get_sobel_horizontal_kernel() {
    struct Kernel kernel = {.width = 3, .height = 3};
    kernel.data[0][0] = -1;
    kernel.data[0][1] = 0;
    kernel.data[0][2] = 1;
    kernel.data[1][0] = -2;
    kernel.data[1][1] = 0;
    kernel.data[1][2] = 2;
    kernel.data[2][0] = -1;
    kernel.data[2][1] = 0;
    kernel.data[2][2] = 1;
    return kernel;
}

static uint8_t get_normalized_channel_value(double value) {
    if (value < 0) {
        value = fabs(value);
    }
    if (value > 255) {
        value = 255;
    }
    return value;
}

struct Bitmap24Pixel get_apply_kernel_pixel_result(const struct Kernel* kernel, const struct Bitmap24Pixel* const bitmap24_input_pixels, size_t row, size_t col,
                                                                int32_t input_width_px, int32_t input_height_px) {
    double red = 0;
    double green = 0;
    double blue = 0;


    for (int64_t i = -(kernel->height - 1) / 2, kernel_i = 0 ; i <= kernel->height / 2; i++, kernel_i++) {
        for (int64_t j = -(kernel->width - 1) / 2, kernel_j = 0; j <= kernel->width / 2; j++, kernel_j++) {
            int64_t index_row = row + i;
            int64_t index_col = col + j;
            index_row = index_row < 0 ? 0 : index_row >= input_height_px ? input_height_px - 1 : index_row;
            index_col = index_col < 0 ? 0 : index_col >= input_width_px ? input_width_px - 1 : index_col;
            const struct Bitmap24Pixel currentPixel = bitmap24_input_pixels[index_row * input_width_px + index_col];

            red += kernel->data[kernel_i][kernel_j] * (double)currentPixel.red;
            green += kernel->data[kernel_i][kernel_j] * (double)currentPixel.green;
            blue += kernel->data[kernel_i][kernel_j] * (double)currentPixel.blue;
        }
    }
    const double divider = kernel->width * kernel->height;
    return (struct Bitmap24Pixel){.red = get_normalized_channel_value(red / divider),
        .green = get_normalized_channel_value(green / divider),
        .blue = get_normalized_channel_value(blue / divider)};
}

struct Bitmap24Pixel get_bitmap24_pixel_rms(const struct Bitmap24Pixel* const firstPixel, const struct Bitmap24Pixel* const secondPixel) {
    return (struct Bitmap24Pixel){.red = (uint8_t)sqrt((double)firstPixel->red * (double)firstPixel->red + (double)secondPixel->red * (double)secondPixel->red),
                                  .green = (uint8_t)sqrt((double)firstPixel->green * (double)firstPixel->green + (double)secondPixel->green * (double)secondPixel->green),
                                  .blue = (uint8_t)sqrt((double)firstPixel->blue * (double)firstPixel->blue + (double)secondPixel->blue * (double)secondPixel->blue)};
}

