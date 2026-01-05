package main

var parenthesesOpen = map[rune]int{'(': 1, '[': 2, '{': 3}
var parenthesesClose = map[rune]int{')': -1, ']': -2, '}': -3}

func isOpenParentheses(char rune) bool {
	return parenthesesOpen[char] != 0
}

func isCloseParentheses(char rune) bool {
	return parenthesesClose[char] != 0
}

func getParenthesesCode(char rune) int {
	return parenthesesOpen[char] + parenthesesClose[char]
}

// 0 ms | Beats 100.00%
func isValid(s string) bool {
	var stack []int = make([]int, 0, 104)
	for _, char := range s {
		if isOpenParentheses(char) {
			stack = append(stack, getParenthesesCode(char))
		} else if isCloseParentheses(char) {
			if len(stack) > 0 {
				topParenthesesCode := stack[len(stack)-1]
				if topParenthesesCode+getParenthesesCode(char) == 0 {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
