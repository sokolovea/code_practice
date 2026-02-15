#pragma once
#include <inttypes.h>
#include <stddef.h>

#pragma pack(push, 1)
struct BitmapFileHeader
{
    uint16_t bfType;
    uint32_t bfSize;
    uint16_t bfReserved1;
    uint16_t bfReserved2;
    uint32_t bfOffBits;
};
#pragma pack(pop)

struct BitmapInfoHeaderV3
{
    uint32_t biSize;
    int32_t biWidth;
    int32_t biHeight;
    uint16_t biPlanes;
    uint16_t biBitCount;
    uint32_t biCompression;
    uint32_t biSizeImage;
    int32_t biXPelsPerMeter;
    int32_t biYPelsPerMeter;
    uint32_t biClrUsed;
    uint32_t biClrImportant;
};

#pragma pack(push, 1)
struct Bitmap24Pixel
{
    uint8_t blue;
    uint8_t green;
    uint8_t red;
};
#pragma pack(pop)


struct Bitmap24Image
{
    size_t width;
    size_t height;
    struct BitmapFileHeader bitmap_file_header;
    struct BitmapInfoHeaderV3 bitmap_info_header_v3;
    struct Bitmap24Pixel* pixels;
};

size_t get_row_size_with_padding(size_t rowSizeWithoutPadding);

int64_t divide_with_ceil(int64_t a, int64_t b);

struct Bitmap24Image get_initialized_bitmap24_image(size_t width, size_t height, struct Bitmap24Pixel* pixels);
