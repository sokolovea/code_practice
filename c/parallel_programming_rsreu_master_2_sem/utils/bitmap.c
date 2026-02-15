#include "bitmap.h"

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
