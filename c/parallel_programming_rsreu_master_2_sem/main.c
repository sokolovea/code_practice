#include <stdio.h>
#include <stdlib.h>
#include "utils/bitmap.h"
#include "utils/kernel.h"

struct Bitmap24Pixel getApplyKernelPixelResult(const struct Kernel* kernel, const struct Bitmap24Pixel* bitmap24_pixel, int64_t row, int64_t col,
                                               int32_t input_width_px, int32_t input_height_px);

static uint8_t getNormalizedChannelValue(double value) {
    if (value < 0) {
        value = abs(value);
    }
    if (value > 255) {
        value = 255;
    }
    return value;
}

struct Bitmap24Pixel getApplyKernelPixelResult(const struct Kernel* kernel, const struct Bitmap24Pixel* const bitmap24_pixel, int64_t row, int64_t col,
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
            const struct Bitmap24Pixel currentPixel = bitmap24_pixel[index_row * input_width_px + index_col];

            red += kernel->data[kernel_i][kernel_j] * (double)currentPixel.red;
            green += kernel->data[kernel_i][kernel_j] * (double)currentPixel.green;
            blue += kernel->data[kernel_i][kernel_j] * (double)currentPixel.blue;
        }
    }
    double divider = kernel->width * kernel->height;
    return (struct Bitmap24Pixel){.red = getNormalizedChannelValue(red / divider),
        .green = getNormalizedChannelValue(green / divider),
        .blue = getNormalizedChannelValue(blue / divider)};
}

int main(int argc, char *argv[]) {
    // struct Bitmap24Image bitmap24Image
    if (argc < 3) {
        fprintf(stderr, "Usage: %s [INPUT_FILE] [OUTPUT_FILE]\n", argv[0]);
        return 1;
    }
    FILE* inputImageFile = fopen(argv[1], "rb");
    FILE* outputImageFile = fopen(argv[2], "wb");

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
        return 3;
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

    const int32_t output_width_px = input_width_px; //DEBUG
    const int32_t output_height_px = input_height_px; //GEBUG

    const uint64_t input_row_bytes_count_without_padding = input_width_px * sizeof(struct Bitmap24Pixel);
    const uint64_t output_row_bytes_count_without_padding = output_width_px * sizeof(struct Bitmap24Pixel);

    const uint64_t input_row_bytes_count_with_padding = get_row_size_with_padding(input_row_bytes_count_without_padding);
    const uint64_t output_row_bytes_count_with_padding = get_row_size_with_padding(output_row_bytes_count_without_padding);

    uint64_t padding_bytes = input_row_bytes_count_with_padding - input_row_bytes_count_without_padding;

    uint8_t* const input_row = malloc(input_row_bytes_count_with_padding);
    uint8_t* const output_row = malloc(output_row_bytes_count_with_padding);
    if (input_row == nullptr || output_row == nullptr) {
        fprintf(stderr, "Not enough memory available!\n");
    }

    struct Bitmap24Pixel* const input_image_pixels = malloc(sizeof(struct Bitmap24Pixel) * input_height_px * input_width_px);

    for (size_t row = 0; row < input_height_px; row++) {
        fread(input_image_pixels + row * input_width_px, sizeof(struct Bitmap24Pixel), input_width_px, inputImageFile);
        fseek(inputImageFile, padding_bytes, SEEK_CUR);
    }


    struct Bitmap24Pixel* const output_image_pixels = malloc(sizeof(struct Bitmap24Pixel) * output_height_px * output_width_px);

    struct Kernel kernel = get_sobel_horizontal_kernel();

    for (size_t row = 0; row < output_height_px; row++) {
        for (size_t col = 0; col < output_width_px; col++) {
            output_image_pixels[row * output_width_px + col] = getApplyKernelPixelResult(&kernel, input_image_pixels, row, col, input_width_px, input_height_px);
        }
    }

    struct Bitmap24Image bitmap24_output_image = get_initialized_bitmap24_image(output_width_px, output_height_px, output_image_pixels);

    fwrite(&(bitmap24_output_image.bitmapFileHeader), sizeof(struct BitmapFileHeader), 1, outputImageFile);
    fwrite(&(bitmap24_output_image.bitmapInfoHeaderV3), sizeof(struct BitmapInfoHeaderV3), 1, outputImageFile);
    char padding_junk[4];
    for (size_t row = 0; row < bitmap24_output_image.height; row++) {
        fwrite(output_image_pixels + row * output_width_px, sizeof(struct Bitmap24Pixel), bitmap24_output_image.width, outputImageFile);
        fwrite(padding_junk, 1, padding_bytes, outputImageFile);
    }

    fclose(outputImageFile);

    free(input_row);
    free(output_row);
    free(input_image_pixels);
    free(output_image_pixels);
    return 0;
}
