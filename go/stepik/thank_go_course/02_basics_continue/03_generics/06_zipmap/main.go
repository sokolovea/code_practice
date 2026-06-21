package main

import "fmt"

// начало решения

// ZipMap возвращает карту, где ключи - элементы из keys, а значения - из vals.
func ZipMap[Key comparable, Val any](keys []Key, vals []Val) map[Key]Val {
	resultMap := make(map[Key]Val)
	lengthMap := min(len(keys), len(vals))
	for i := range lengthMap {
		resultMap[keys[i]] = vals[i]
	}
	return resultMap
}

// конец решения

func main() {
	m1 := ZipMap([]string{"one", "two", "thr"}, []int{11, 22, 33})
	fmt.Println(m1)
	// map[one:11 two:22 thr:33]

	m2 := ZipMap([]string{"one"}, []int{11, 22, 33})
	fmt.Println(m2)
	// map[one:11]

	m3 := ZipMap([]string{"one", "two", "thr"}, []int{11})
	fmt.Println(m3)
	// map[one:11]

	m4 := ZipMap([]string{}, []int{11, 22, 33})
	fmt.Println(m4)
	// map[]
}
