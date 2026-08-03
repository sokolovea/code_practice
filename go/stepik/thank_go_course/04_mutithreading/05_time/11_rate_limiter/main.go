// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	interval := time.Second / time.Duration(limit)
	ticker := time.NewTicker(interval)
	cancelChan := make(chan struct{}, 1)
	ticker.Reset(interval)
	handle = func() error {
		select {
		case <-ticker.C:
			go func() {
				fn()
			}()
		case <-cancelChan:
			return ErrCanceled
		}
		return nil
	}
	isCancelled := false
	cancel = func() {
		if !isCancelled {
			isCancelled = true
			close(cancelChan)
			ticker.Stop()
		}
	}
	return
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for range n {
		handle()
	}
	cancel()
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
