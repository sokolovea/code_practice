package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	cancelChan := make(chan struct{}, 1)
	cancelFunc := func() {
		select {
		case cancelChan <- struct{}{}:
		default:
			return
		}
	}
	go func() {
		ticker := time.NewTicker(dur)
		for {
			select {
			case <-cancelChan:
				ticker.Stop()
				return
			case <-ticker.C:
				fn()
			}
		}
	}()
	return cancelFunc
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
