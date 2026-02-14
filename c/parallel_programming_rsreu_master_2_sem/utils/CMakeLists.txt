cmake_minimum_required(VERSION 4.0)
project(utils C)

set(CMAKE_C_STANDARD 23)

add_library(bmp STATIC
        bitmap.h
        bitmap.c)
