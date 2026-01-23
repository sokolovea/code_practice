package main

import (
	"sync"
)

type Foo struct {
	second sync.WaitGroup
	third  sync.WaitGroup
}

func NewFoo() *Foo {
	var foo Foo
	foo.second.Add(1)
	foo.third.Add(1)
	return &foo
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.second.Done()
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	f.second.Wait()
	printSecond()
	f.third.Done()
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	f.third.Wait()
	printThird()
}
