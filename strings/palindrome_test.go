package strings

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("nawan") {
		t.Error("not a palindrome")
	}
	if !IsPalindrome("nawwan") {
		t.Error("not a palindrome")
	}
	if !IsPalindrome("Tattarrattat") {
		t.Error("not a palindrome")
	}
	if !IsPalindrome("Aibohphobia") {
		t.Error("not a palindrome")
	}
	if !IsPalindromeRecursive("nawan") {
		t.Error("not a palindrome")
	}
	if !IsPalindromeRecursive("nawwan") {
		t.Error("not a palindrome")
	}
	if !IsPalindromeRecursive("Tattarrattat") {
		t.Error("not a palindrome")
	}
	if !IsPalindromeRecursive("Aibohphobia") {
		t.Error("not a palindrome")
	}
}
