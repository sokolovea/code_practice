package majorityelement

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

var testData = []struct {
	inputSlice []int
	result     int
}{
	{[]int{0}, 0},
	{[]int{1}, 1},
	{[]int{1, 1, 1}, 1},
	{[]int{1, 2, 3, 2, 2, 2, 7, 2, 2, 9}, 2},
	{[]int{1, 9, 9, 4, 2, 9, 9, 9, 9, 9, 0, 9, 10}, 9},
	{[]int{9, -1, 7, 0, -1, -1, -1, -1, -1, 12, -1, -1, 4}, -1},
}

func majorityElementInnerTest(t *testing.T, f func([]int) int) {
	for i, test := range testData {
		result := f(test.inputSlice)
		if result != test.result {
			t.Errorf("%d) %v: [result, expected] = [%v; %v]",
				i, test.inputSlice, result, test.result)
		}
	}
}

func majorityElementInnerBench(b *testing.B, f func([]int) int) {
	benchCases := make([][]int, 7)
	for i := range benchCases {
		benchCases[i] = make([]int, int(math.Pow10(i)))
		for j := range benchCases[i] {
			//only bor bench, no majority element warranties
			benchCases[i][j] = rand.Intn(math.MaxInt16)
		}
	}

	for _, testCase := range benchCases {
		b.Run(fmt.Sprintf("Bench on %d elements", len(testCase)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = f(testCase)
				}
			})
	}
}

func TestMajorityElementMap(t *testing.T) {
	majorityElementInnerTest(t, majorityElementMap)
}

func TestMajorityElementSort(t *testing.T) {
	majorityElementInnerTest(t, majorityElementSort)
}

func TestMajorityElementBoyerMoore(t *testing.T) {
	majorityElementInnerTest(t, majorityElementBoyerMoore)
}

func BenchmarkMajorityElementMap(b *testing.B) {
	majorityElementInnerBench(b, majorityElementMap)
}

func BenchmarkMajorityElementSort(b *testing.B) {
	majorityElementInnerBench(b, majorityElementSort)
}

func BenchmarkMajorityElementBoyerMoore(b *testing.B) {
	majorityElementInnerBench(b, majorityElementBoyerMoore)
}
