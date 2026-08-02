package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancelChan := make(chan struct{}, 1)
	cancelFunc := func() {
		select {
		case cancelChan <- struct{}{}:
		default:
			return
		}
	}
	go func() {
		timer := time.NewTimer(dur)
		select {
		case <-cancelChan:
		case <-timer.C:
			fn()
		}
	}()
	return cancelFunc
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
