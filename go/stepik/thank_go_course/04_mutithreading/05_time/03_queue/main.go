package main

import (
	"errors"
	"time"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue chan int

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if block {
		return <-q, nil
	} else {
		select {
		case res := <-q:
			return res, nil
		default:
			return 0, ErrEmpty
		}
	}
}

// Put помещает элемент в очередь.
// Если очередь заполнена и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if block {
		q <- val
		return nil
	} else {
		select {
		case q <- val:
			return nil
		default:
			return ErrFull
		}
	}
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return make(chan int, n)
}

// конец решения

func main() {
	q := MakeQueue(1)
	q.Put(1, true)

	done := make(chan struct{})
	go func() {
		q.Put(2, true)
		close(done)
	}()

	time.Sleep(10 * time.Millisecond)
	q.Get(true)
	<-done
}
