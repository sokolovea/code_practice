#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <omp.h>
#include "utils/bitmap.h"
#include "utils/kernel.h"


int main(int argc, char *argv[]) {
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

    if (inputImageFile == nullptr || outputImageFile == nullptr) {
        fprintf(stderr, "Can't open image files!\n");
        return 2;
    }
    struct BitmapFileHeader bitmapInputFileHeader;
    struct BitmapInfoHeaderV3 bitmapInputInfoHeaderV3;
    size_t count_handled_objects = 0;

    count_handled_objects = fread(&bitmapInputFileHeader, sizeof(struct BitmapFileHeader), 1, inputImageFile);
    if (count_handled_objects == 0 || bitmapInputFileHeader.bfType != 0x4d42) {
        fprintf(stderr, "Not BMP file!\n");
        return 3;
    }

    count_handled_objects = fread(&bitmapInputInfoHeaderV3, sizeof(struct BitmapInfoHeaderV3), 1, inputImageFile);
    if (count_handled_objects == 0) {
        fprintf(stderr, "Not valid structure of BMP file\n");
        return 4;
    }

    int fseek_result = 0;
    fseek_result = fseek(inputImageFile, bitmapInputFileHeader.bfOffBits, SEEK_SET);

    if (fseek_result != 0) {
        fprintf(stderr, "Not valid structure of BMP file\n");
        return 4;
    }

    if (bitmapInputInfoHeaderV3.biBitCount != 24) {
        fprintf(stderr, "Not 24 bit BMP file!\n");
        return 5;
    }

    const int32_t input_width_px = bitmapInputInfoHeaderV3.biWidth;
    const int32_t input_height_px = bitmapInputInfoHeaderV3.biHeight;

    const int32_t output_width_px = input_width_px;
    const int32_t output_height_px = input_height_px;

    const size_t input_row_bytes_count_without_padding = input_width_px * sizeof(struct Bitmap24Pixel);
    const size_t output_row_bytes_count_without_padding = output_width_px * sizeof(struct Bitmap24Pixel);

    const size_t input_row_bytes_count_with_padding = get_row_size_with_padding(input_row_bytes_count_without_padding);
    const size_t output_row_bytes_count_with_padding = get_row_size_with_padding(output_row_bytes_count_without_padding);

    const size_t padding_bytes = input_row_bytes_count_with_padding - input_row_bytes_count_without_padding;

    uint8_t* const input_row = malloc(input_row_bytes_count_with_padding);
    uint8_t* const output_row = malloc(output_row_bytes_count_with_padding);
    if (input_row == nullptr || output_row == nullptr) {
        fprintf(stderr, "Not enough memory available!\n");
        return 6;
    }

    struct Bitmap24Pixel* const input_image_pixels = malloc(sizeof(struct Bitmap24Pixel) * input_height_px * input_width_px);

    for (size_t row = 0; row < input_height_px; row++) {
        count_handled_objects = fread(input_image_pixels + row * input_width_px, sizeof(struct Bitmap24Pixel), input_width_px, inputImageFile);
        fseek_result = fseek(inputImageFile, padding_bytes, SEEK_CUR);

        if (count_handled_objects != input_width_px || fseek_result != 0) {
            fprintf(stderr, "Not valid bitmap data of BMP file\n");
            return 4;
        }
    }


    struct Bitmap24Pixel* const output_image_pixels = malloc(sizeof(struct Bitmap24Pixel) * output_height_px * output_width_px);

    struct Kernel kernel_first = get_sobel_horizontal_kernel();
    struct Kernel kernel_second = get_sobel_vertical_kernel();

    printf("Started calculations: iterations count = %"PRId64"\n", iterations_count);

    struct timespec start, end;

#if 1
    clock_gettime(CLOCK_REALTIME, &start);
    for (size_t counter = 0; counter < iterations_count; counter++) {
        for (size_t row = 0; row < output_height_px; row++) {
            for (size_t col = 0; col < output_width_px; col++) {
                auto pixel_by_first_kernel = get_apply_kernel_pixel_result(&kernel_first, input_image_pixels, row, col, input_width_px, input_height_px);
                auto pixel_by_second_kernel = get_apply_kernel_pixel_result(&kernel_second, input_image_pixels, row, col, input_width_px, input_height_px);
                output_image_pixels[row * output_width_px + col] = get_bitmap24_pixel_rms(&pixel_by_first_kernel, &pixel_by_second_kernel);
            }
        }
    }
    clock_gettime(CLOCK_REALTIME, &end);

    double elapsed_s = (end.tv_sec - start.tv_sec);
    elapsed_s += (end.tv_nsec - start.tv_nsec) / 1'000'000'000.0;
    printf("Done!\n");
    printf("Elapsed time: %.3lf s for all iterations;\n", elapsed_s);
    printf("              %.3lf s for one iterations.\n", elapsed_s / iterations_count);
#endif

#if 0
    omp_set_num_threads(8);
    int threads_count = omp_get_max_threads();
    clock_gettime(CLOCK_REALTIME, &start);
    for (size_t counter = 0; counter < iterations_count; counter++) {
#pragma omp parallel for
        for (size_t row = 0; row < output_height_px; row++) {
            for (size_t col = 0; col < output_width_px; col++) {
                auto pixel_by_first_kernel = get_apply_kernel_pixel_result(&kernel_first, input_image_pixels, row, col, input_width_px, input_height_px);
                auto pixel_by_second_kernel = get_apply_kernel_pixel_result(&kernel_second, input_image_pixels, row, col, input_width_px, input_height_px);
                output_image_pixels[row * output_width_px + col] = get_bitmap24_pixel_rms(&pixel_by_first_kernel, &pixel_by_second_kernel);
            }
        }
    }
    clock_gettime(CLOCK_REALTIME, &end);

    double elapsed_s = (end.tv_sec - start.tv_sec);
    elapsed_s += (end.tv_nsec - start.tv_nsec) / 1'000'000'000.0;
    printf("Done! Threads count = %d.\n", threads_count);
    printf("Elapsed time: %.3lf s for all iterations;\n", elapsed_s);
    printf("              %.3lf s for one iterations.\n", elapsed_s / iterations_count);
#endif
    struct Bitmap24Image bitmap24_output_image = get_initialized_bitmap24_image(output_width_px, output_height_px, output_image_pixels);

    size_t fwrite_result = 0;
    fwrite_result = fwrite(&(bitmap24_output_image.bitmap_file_header), sizeof(struct BitmapFileHeader), 1, outputImageFile);
    if (fwrite_result != 1) {
        fprintf(stderr, "Cannot write bitmap_file_header to output file!\n");
        return 7;
    }
    fwrite_result = fwrite(&(bitmap24_output_image.bitmap_info_header_v3), sizeof(struct BitmapInfoHeaderV3), 1, outputImageFile);
    if (fwrite_result != 1) {
        fprintf(stderr, "Cannot write bitmap_info_header_v3 to output file!\n");
        return 7;
    }
    char padding_junk[4];
    for (size_t row = 0; row < bitmap24_output_image.height; row++) {
        fwrite_result = fwrite(output_image_pixels + row * output_width_px, sizeof(struct Bitmap24Pixel), bitmap24_output_image.width, outputImageFile);
        if (fwrite_result != bitmap24_output_image.width) {
            fprintf(stderr, "Cannot write bitmap data to output file!\n");
            return 7;
        }
        fwrite_result = fwrite(padding_junk, sizeof(char), padding_bytes, outputImageFile);
        if (fwrite_result != padding_bytes) {
            fprintf(stderr, "Cannot write bitmap data to output file!\n");
            return 7;
        }
    }

    fclose(outputImageFile);

    free(input_row);
    free(output_row);
    free(input_image_pixels);
    free(output_image_pixels);
    return 0;
}
