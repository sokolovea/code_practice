package lab8

import (
	"fmt"
	"testing"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BenchmarkBruteForceDummy(b *testing.B) {
	for _, eps := range c.EpsilonsTest {
		b.Run(fmt.Sprintf("eps=%v", eps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BruteForceDummy(-5.0, 5.0, -5.0, 5.0, c.FTest, eps)
			}
		})
	}
}
