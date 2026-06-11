/*
 * Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез.
 * Функция-предикат вызывается для каждого элемента исходного среза.
 * Если она возвращает true, элемент попадает в отфильтрованный срез. Если возвращает false — не попадает.
 * Считайте исходный срез из стандартного ввода с помощью готовой функции readInput().
 * Затем выполните на нем filter(). В качестве предиката используйте функцию, которая возвращает true только для четных чисел.
 * Гарантируется, что на вход подаются только целые числа. Напечатайте отфильтрованный срез.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int = make([]int, 0, len(iterable))
	for _, number := range iterable {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	src := readInput()
	res := filter(func(number int) bool { return number%2 == 0 }, src)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
