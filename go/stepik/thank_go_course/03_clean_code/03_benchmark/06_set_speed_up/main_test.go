package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkIntSet(b *testing.B) {
	for setSize := range []int{10, 100, 1_000, 10_000, 100_000, 1_000_000} {
		intSet := fillIntSet(setSize)
		b.Run(fmt.Sprintf("%d", setSize), func(b *testing.B) {
			for b.Loop() {
				elem := rand.Intn(100_000)
				intSet.Contains(elem)
			}
		})
	}
}

func fillIntSet(size int) IntSet {
	intSet := MakeIntSet()
	for range size {
		elem := rand.Intn(100_000)
		intSet.Add(elem)
	}
	return intSet
}
