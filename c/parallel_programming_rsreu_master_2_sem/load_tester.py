#!/usr/bin/env python3
"""
Скрипт для нагрузочного тестирования OpenMP-приложения.
Читает конфигурацию из CSV, запускает программу с различными параметрами
привязки потоков и сохраняет результаты в выходной CSV.
"""

import argparse
import csv
import os
import re
import subprocess
import sys
import time
from typing import Optional


def parse_arguments():
    parser = argparse.ArgumentParser(
        description="Нагрузочное тестирование OpenMP-приложения."
    )
    parser.add_argument(
        "--csv",
        required=True,
        help="Путь к входному CSV-файлу с конфигурацией тестов."
    )
    parser.add_argument(
        "--exe",
        default="bin/a.out",
        help="Путь к исполняемому файлу программы (по умолчанию bin/a.out)."
    )
    parser.add_argument(
        "--output",
        default="results.csv",
        help="Файл для записи результатов (по умолчанию results.csv)."
    )
    parser.add_argument(
        "--delay",
        type=int,
        default=10,
        help="Задержка между запусками в секундах (по умолчанию 10)."
    )
    return parser.parse_args()


def setup_openmp_env(num_threads: int, places: str, proc_bind: str):
    """Устанавливает переменные окружения для управления нитями OpenMP."""
    os.environ["OMP_NUM_THREADS"] = str(num_threads)
    if places:
        os.environ["OMP_PLACES"] = places
    elif "OMP_PLACES" in os.environ:
        del os.environ["OMP_PLACES"]
    if proc_bind:
        os.environ["OMP_PROC_BIND"] = proc_bind
    elif "OMP_PROC_BIND" in os.environ:
        del os.environ["OMP_PROC_BIND"]


def extract_elapsed_time(stdout: str) -> Optional[float]:
    """
    Извлекает суммарное время выполнения из stdout программы.
    Ожидается строка вида: "Elapsed time: 12.345 s for all iterations;"
    """
    pattern = r"Elapsed time:\s+([0-9.]+)\s+s for all iterations;"
    match = re.search(pattern, stdout)
    if match:
        return float(match.group(1))
    return None


def run_single_test(
    exe_path: str,
    input_file: str,
    output_file: str,
    iterations: int,
    num_threads: int,
    places: str,
    proc_bind: str,
    test_name: str,
    run_number: int,
    delay: int,
    writer,
    out_csv,
):
    """Запускает один тест, извлекает время и записывает результат."""
    print(f"[{test_name}, run {run_number}] Запуск...")
    setup_openmp_env(num_threads, places, proc_bind)

    try:
        result = subprocess.run(
            [exe_path, input_file, output_file, str(iterations)],
            capture_output=True,
            text=True,
            check=False,
        )
    except FileNotFoundError:
        print(
            f"Ошибка: исполняемый файл {exe_path} не найден.", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Ошибка при запуске: {e}", file=sys.stderr)
        writer.writerow([test_name, run_number, "ERROR"])
        out_csv.flush()
        time.sleep(delay)
        return

    if result.returncode != 0:
        print(f"Программа завершилась с ошибкой (код {result.returncode})")
        print("STDERR:", result.stderr)
        writer.writerow([test_name, run_number, "ERROR"])
        out_csv.flush()
        time.sleep(delay)
        return

    elapsed = extract_elapsed_time(result.stdout)
    if elapsed is None:
        print("Не удалось извлечь время из вывода программы.")
        print("STDOUT:", result.stdout)
        writer.writerow([test_name, run_number, "PARSE_ERROR"])
    else:
        print(f"Время выполнения: {elapsed:.3f} с")
        writer.writerow([test_name, run_number, f"{elapsed:.3f}"])

    out_csv.flush()
    time.sleep(delay)


def main():
    args = parse_arguments()

    if not os.path.isfile(args.csv):
        print(f"Ошибка: файл {args.csv} не найден.", file=sys.stderr)
        sys.exit(1)

    # Чтение конфигурации
    with open(args.csv, "r", newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        required_fields = {"test_name", "runs", "input_file",
                           "iterations", "num_threads", "places", "proc_bind"}
        if not required_fields.issubset(reader.fieldnames):
            missing = required_fields - set(reader.fieldnames)
            print(
                f"Ошибка: в CSV отсутствуют необходимые поля: {missing}", file=sys.stderr)
            sys.exit(1)

        tests = list(reader)

    with open(args.output, "w", newline="", encoding="utf-8") as out_csv:
        writer = csv.writer(out_csv)
        writer.writerow(["test_name", "run", "elapsed_time_s"])

        for test in tests:
            test_name = test["test_name"].strip()
            print("DEBUG: ", test_name)
            runs = int(test["runs"])
            input_file = test["input_file"].strip()
            print("DEBUG: ", input_file)
            iterations = int(test["iterations"])
            num_threads = int(test["num_threads"])
            places = test["places"].strip() if test["places"] else ""
            proc_bind = test["proc_bind"].strip() if test["proc_bind"] else ""

            for run in range(1, runs + 1):
                output_image = f"output_{test_name}_{run}.bmp"

                run_single_test(
                    exe_path=args.exe,
                    input_file=input_file,
                    output_file=output_image,
                    iterations=iterations,
                    num_threads=num_threads,
                    places=places,
                    proc_bind=proc_bind,
                    test_name=test_name,
                    run_number=run,
                    delay=args.delay,
                    writer=writer,
                    out_csv=out_csv,
                )

    print(f"Результаты сохранены в {args.output}")


if __name__ == "__main__":
    main()
