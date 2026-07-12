package main

// не удаляйте импорты, они используются при проверке

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
type IntSet struct {
	data map[int]struct{}
}

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	return IntSet{data: make(map[int]struct{})}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
	_, isContains := s.data[elem]
	return isContains
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s.data[elem] = struct{}{}
	return true
}
