package reorderlist

import "testing"

var testData = []struct {
	inputSlice       []int
	inputStringRepr  string
	outputStringRepr string
}{
	{[]int{}, "", ""},
	{[]int{1}, "[1]", "[1]"},
	{[]int{2}, "[2]", "[2]"},
	// {[]int{1, 2}, "[1->2]", "[2->1]"}, //bugs
	{[]int{1, 2, 3}, "[1]->[2]->[3]", "[1]->[3]->[2]"},
	{[]int{1, 2, 3, 4}, "[1]->[2]->[3]->[4]", "[1]->[4]->[2]->[3]"},
	{[]int{1, 2, 3, 4, 5}, "[1]->[2]->[3]->[4]->[5]", "[1]->[5]->[2]->[4]->[3]"},
	{[]int{1, 2, 3, 4, 5, 6}, "[1]->[2]->[3]->[4]->[5]->[6]", "[1]->[6]->[2]->[5]->[3]->[4]"},
}

func sliceToList(inputSlice []int) *ListNode {
	var firstElement *ListNode = nil
	root := &firstElement
	currentElement := root
	var prevElement *ListNode = nil
	for _, value := range inputSlice {
		*currentElement = new(ListNode)
		(**currentElement).Val = value
		(**currentElement).Next = nil
		if prevElement != nil {
			(*prevElement).Next = *currentElement
		}
		prevElement = *currentElement
		currentElement = &((**currentElement).Next)
	}
	return *root
}

func TestSliceToList(t *testing.T) {
	for i, currentTestData := range testData {
		got := sliceToList(currentTestData.inputSlice).String()
		expected := currentTestData.inputStringRepr
		if got != expected {
			t.Errorf("Error test case %d: expected = %s; got = %s", i+1, expected, got)
		}
	}
}

func TestReorderList(t *testing.T) {
	for i, currentTestData := range testData {
		head := sliceToList(currentTestData.inputSlice)
		reorderList(head)
		got := head.String()
		expected := currentTestData.outputStringRepr
		if got != expected {
			t.Errorf("Error test case %d: expected = %s; got = %s", i+1, expected, got)
		}
	}
}
