cmake_minimum_required(VERSION 4.0)
project(parallel_programming_rsreu_master_2_sem C)

set(CMAKE_C_STANDARD 23)

add_subdirectory(utils)
add_executable(parallel_programming_rsreu_master_2_sem main.c
        utils/kernel.h
)

target_link_libraries(parallel_programming_rsreu_master_2_sem
        utils
)

target_include_directories(parallel_programming_rsreu_master_2_sem
        PUBLIC
        ${CMAKE_CURRENT_SOURCE_DIR}/utils
)