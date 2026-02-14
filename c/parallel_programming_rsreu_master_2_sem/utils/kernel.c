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