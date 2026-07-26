package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue struct {
	buffer []int
	tail   int
	head   int
	length int
	tokens chan struct{}
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q *Queue) Get(block bool) (int, error) {
	if q.tail == q.head {
		if block {
			<-q.tokens
			q.tokens <- struct{}{}
		} else {
			return 0, ErrEmpty
		}
	}
	result := q.buffer[q.head]
	q.head = (q.head + 1) % (q.length + 1)
	<-q.tokens
	return result, nil
}

// Put помещает элемент в очередь.
// Если очередь заполнена и block = false -
// возвращает ошибку.
func (q *Queue) Put(val int, block bool) error {
	if q.tail == q.length && q.head == 0 || q.head == q.tail+1 {
		if block {
			q.tokens <- struct{}{}
			<-q.tokens
		} else {
			return ErrFull
		}
	}
	q.buffer[q.tail] = val
	q.tail = (q.tail + 1) % (q.length + 1)
	q.tokens <- struct{}{}
	return nil
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	q := Queue{buffer: make([]int, n+1), length: n, tokens: make(chan struct{}, n)}
	return q
}

// конец решения

func main() {
	q := MakeQueue(1)

	go func() {
		err := q.Put(11, true)
		fmt.Println("put 11:", err)
		// put 11: <nil>

		err = q.Put(12, true)
		fmt.Println("put 12:", err)
		// put 12: <nil>
	}()
	res, err := q.Get(true)
	fmt.Println("get:", res, err)
	// get: 11 <nil>

	res, err = q.Get(true)
	fmt.Println("get:", res, err)
	// get: 12 <nil>
}
