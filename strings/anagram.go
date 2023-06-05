package strings

import (
	strs "strings"
)

// Problem: Given two strings, check whether two strings are an anagram
// of each other. Two strings are said to be an anagram of each other if
// they are just permutations of each other. That is, the set of characters
// in both the strings must be the same, only the order of characters can
// be different.

func IsAnagram(str0, str1 string) bool {
	if len(str0) != len(str1) {
		return false
	}

	str0, str1 = strs.ToLower(str0), strs.ToLower(str1)

	charMap := make(map[byte]uint)

	for i := 0; i < len(str0); i++ {
		charMap[str0[i]]++
	}

	for i := 0; i < len(str1); i++ {
		charMap[str1[i]]++
	}

	for _, v := range charMap {
		if v != 2 {
			return false
		}
	}

	return true
}
