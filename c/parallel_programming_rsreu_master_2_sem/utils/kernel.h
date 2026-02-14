#pragma once

#include <stddef.h>
#include <stdint.h>

#define MAX_KERNEL_SIZE 50

struct Kernel {
    double data[MAX_KERNEL_SIZE][MAX_KERNEL_SIZE];
    int64_t width;
    int64_t height;
};

struct Kernel get_sobel_vertical_kernel();

struct Kernel get_sobel_horizontal_kernel();