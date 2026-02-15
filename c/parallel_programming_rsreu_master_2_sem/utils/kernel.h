#pragma once

#include <math.h>
#include <stdint.h>
#include "bitmap.h"

#define MAX_KERNEL_SIZE 50

struct Kernel {
    double data[MAX_KERNEL_SIZE][MAX_KERNEL_SIZE];
    int64_t width;
    int64_t height;
};

struct Kernel get_sobel_vertical_kernel();

struct Kernel get_sobel_horizontal_kernel();

struct Bitmap24Pixel get_apply_kernel_pixel_result(const struct Kernel* kernel, const struct Bitmap24Pixel* bitmap24_input_pixels,  size_t row,  size_t col,
                                               int32_t input_width_px, int32_t input_height_px);

struct Bitmap24Pixel get_bitmap24_pixel_rms(const struct Bitmap24Pixel* const firstPixel, const struct Bitmap24Pixel* const secondPixel);