// Promise.all()
package main

import (
	"fmt"
	"time"
)

// начало решения

type pair struct {
	data  any
	order int
}

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// Выполните все переданные функции,
	// соберите результаты в срез и верните его.
	outChan := make(chan pair, len(funcs))
	done := make(chan struct{})
	go func() {
		counter := 0
		for range done {
			counter++
			if counter == len(funcs) {
				close(outChan)
				close(done)
				return
			}
		}
	}()
	for i, f := range funcs {
		go func() {
			outChan <- pair{f(), i}
			done <- struct{}{}
		}()
	}
	resultSlise := make([]any, len(funcs))
	for res := range outChan {
		resultSlise[res.order] = res.data
	}
	return resultSlise
}

// конец решения

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
