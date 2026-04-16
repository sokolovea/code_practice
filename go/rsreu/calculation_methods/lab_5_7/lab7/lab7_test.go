package lab7

import (
	"fmt"
	"testing"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BenchmarkBisection(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Bisection(0.0, 2.0, c.FTest, c.GTest, eps)
			}
		})
	}
}

func BenchmarkChord(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Chord(0.0, 2.0, c.FTest, c.GTest, eps)
			}
		})
	}
}

func BenchmarkNewton(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Newton(0.0, 2.0, c.FTest, c.GTest, c.G2Test, eps)
			}
		})
	}
}
