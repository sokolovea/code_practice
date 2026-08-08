package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// начало решения

type Total struct {
	counter atomic.Int32
}

// увеличивает счетчик на 1
func (total *Total) Increment() {
	total.counter.Add(1)
}

// возвращает значение счетчика
func (total *Total) Value() int {
	return int(total.counter.Load())
}

// конец решения

func main() {
	var wg sync.WaitGroup

	var total Total

	for range 5 {
		wg.Go(func() {
			for range 10000 {
				total.Increment()
			}
		})
	}

	wg.Wait()
	fmt.Println("total", total.Value())
}
