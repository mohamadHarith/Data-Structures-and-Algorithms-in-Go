package strings

import "testing"

func TestIsAnagram(t *testing.T) {
	if !IsAnagram("eat", "Tea") {
		t.Error("not an anagram")
	}
	if !IsAnagram("eat", "ate") {
		t.Error("not an anagram")
	}
	if IsAnagram("eat", "sat") {
		t.Error("not an anagram")
	}
}
