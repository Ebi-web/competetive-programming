package main

func longestPalindrome(s string) string {
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindrome(s[j : j+i]) {
				return s[j : j+i]
			}
		}
	}
	return ""
}

// judge whether a string is palindrome
func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
