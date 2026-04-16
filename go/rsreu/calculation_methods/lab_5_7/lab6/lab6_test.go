package lab6

import (
	"fmt"
	"testing"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BenchmarkParabola(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Parabola(0.0, 2.0, c.FTest, eps)
			}
		})
	}
}

func BenchmarkGoldenRatio(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GoldenRatio(0.0, 2.0, c.FTest, eps)
			}
		})
	}
}
