package main

import (
	"io"
)

// TokenReader начитывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
}

// начало решения

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
	tokensCount := 0
	for {
		token, err := src.ReadToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return tokensCount, err
		}
		if predicate(token) {
			err = dst.WriteToken(token)
			if err != nil {
				return tokensCount, err
			}
			tokensCount++
		}
	}
	return tokensCount, nil
}

// // конец решения

// func main() {
// 	// Для проверки придется создать конкретные типы,
// 	// которые реализуют интерфейсы TokenReader и TokenWriter.

// 	// Ниже для примера используются NewWordReader и NewWordWriter,
// 	// но вы можете сделать любые на свое усмотрение.

// 	r := NewWordReader("go is awesome")
// 	w := NewWordWriter()
// 	predicate := func(s string) bool {
// 		return s != "is"
// 	}
// 	n, err := FilterTokens(w, r, predicate)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%d tokens: %v\n", n, w.Words())
// 	// 2 tokens: [go awesome]
// }
