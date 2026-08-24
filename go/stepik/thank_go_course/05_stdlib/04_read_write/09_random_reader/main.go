package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

// начало решения

type RandomBytesReader struct {
	maxBytesCount   int
	currBytesReaded int
}

func (rbr *RandomBytesReader) Read(p []byte) (n int, err error) {
	if rbr.currBytesReaded >= rbr.maxBytesCount {
		return 0, io.EOF
	}
	bytesToFill := min(rbr.maxBytesCount-rbr.currBytesReaded, len(p))
	rbr.currBytesReaded += bytesToFill
	return rand.Read(p[:bytesToFill])
}

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {
	return &RandomBytesReader{maxBytesCount: max, currBytesReaded: 0}
}

// конец решения

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}
