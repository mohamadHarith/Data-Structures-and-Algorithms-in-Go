package strings

import (
	strs "strings"
)

// Problem: Check if a given string is a palindrome.

// IsPalindrome : checks if a string is a palindrome
func IsPalindrome(str string) bool {
	for i := 0; i < (len(str)/2)-1; i++ {
		if !strs.EqualFold(string(str[i]), string(str[len(str)-1-i])) {
			return false
		}
	}

	return true
}

// IsPalindromeRecursive : a recursive approach to check if a string is a palindrome
func IsPalindromeRecursive(str string) bool {
	// base case 1
	if len(str) <= 1 {
		return true
	}

	// base case 2
	if !strs.EqualFold(string(str[0]), string(str[len(str)-1])) {
		return false
	}

	return IsPalindromeRecursive(str[1 : len(str)-1])
}
