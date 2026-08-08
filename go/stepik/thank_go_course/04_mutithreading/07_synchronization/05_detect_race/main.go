package main

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	alive := make(chan struct{}) // (1)
	close(alive)                 // (2)

	go func() {
		time.Sleep(duration)
		select {
		case <-alive: // (3)
			fn()
		default:
		}
	}()

	cancel := func() {
		alive = nil // (4)
	}
	return cancel
}

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(50*time.Millisecond, work)
	time.Sleep(50 * time.Millisecond)
	go cancel()
}
