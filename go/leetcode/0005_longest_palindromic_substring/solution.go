package main

// O(N)
// Only ASCII and case-sensitive
func isPalindrome(s string) bool {
	var result bool = true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			result = false
			break
		}
	}
	return result
}

// O(N^3)
// 409 ms | Beats 7.95%
func longestPalindromeSimple(s string) string {
	maxPalindrome := ""
	maxPalindromeLength := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			currentSubstring := s[i : j+1]
			if isPalindrome(currentSubstring) && len(currentSubstring) > maxPalindromeLength {
				maxPalindrome = currentSubstring
				maxPalindromeLength = len(currentSubstring)
			}
		}
	}
	return maxPalindrome
}

// TODO: O(N) with Manacher’s Algorithm
// func longestPalindromeOptimized(s string) string {
// }
