cmake_minimum_required(VERSION 4.0)
project(parallel_programming_rsreu_master_2_sem C)

set(CMAKE_C_STANDARD 23)

# Сначала находим все зависимости
#find_package(OpenMP REQUIRED)
find_package(MPI REQUIRED)

# Добавляем поддиректорию до создания исполняемого файла
add_subdirectory(utils)

# Создаём исполняемый файл
add_executable(parallel_programming_rsreu_master_2_sem
        main.c
        utils/kernel.h
)

# Устанавливаем флаги оптимизации для конкретной цели
target_compile_options(parallel_programming_rsreu_master_2_sem
        PRIVATE
        $<$<CONFIG:Debug>:-Og -g>
        $<$<CONFIG:Release>:-O2>
)

# Указываем пути включения для цели
target_include_directories(parallel_programming_rsreu_master_2_sem
        PUBLIC
        ${CMAKE_CURRENT_SOURCE_DIR}/utils
        ${MPI_INCLUDE_PATH}
)

# Применяем флаги OpenMP к цели
target_link_libraries(parallel_programming_rsreu_master_2_sem
        PRIVATE
        utils
        m
#        OpenMP::OpenMP_C
        ${MPI_C_LIBRARIES}
)
