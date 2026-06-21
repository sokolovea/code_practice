/*
 * Реализуйте обобщенный тип Map со следующими методами:
 *
 * Set устанавливает значение для ключа.
 * Get возвращает значение по ключу.
 * Keys возвращает срез ключей карты.
 * Values возвращает срез значений карты.
 */

package main

import "fmt"

// начало решения

// Map - карта "ключ-значение".
type Map[Key comparable, Value any] map[Key]Value

// Set устанавливает значение для ключа.
func (m Map[Key, Value]) Set(key Key, val Value) {
	m[key] = val
}

// Get возвращает значение по ключу.
func (m Map[Key, Value]) Get(key Key) Value {
	return m[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m Map[Key, Value]) Keys() []Key {
	keys := make([]Key, 0, len(m))
	for eachKey, _ := range m {
		keys = append(keys, eachKey)
	}
	return keys
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m Map[Key, Value]) Values() []Value {
	vals := make([]Value, 0, len(m))
	for _, eachVal := range m {
		vals = append(vals, eachVal)
	}
	return vals
}

// конец решения

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
