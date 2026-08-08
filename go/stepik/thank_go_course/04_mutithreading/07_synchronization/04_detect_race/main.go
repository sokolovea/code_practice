package main

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	canceled := false // (1)

	go func() {
		time.Sleep(duration)
		if !canceled { // (2)
			fn()
		}
	}()

	cancel := func() {
		canceled = true // (3)
	}
	return cancel // (4)
}

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(50*time.Millisecond, work)
	time.Sleep(50 * time.Millisecond)
	go cancel()
}
