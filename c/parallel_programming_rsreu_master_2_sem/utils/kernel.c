#include <stdlib.h>
#include "kernel.h"

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