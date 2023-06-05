package strings

// Problem: Given a string of words, reverse the words.
// Ex: "Hello, world!" should become "world! Hello,"

func ReverseWordsInString(str string) string {

	var chars = []byte(str)

	start := 0
	for end := 0; end < len(chars); end++ {
		if chars[end] == ' ' {
			ReverseWord(chars, start, end-1)
			start = end + 1
		}
	}

	// reverse the last word
	ReverseWord(chars, start, len(chars)-1)

	// reverse the entire string
	ReverseWord(chars, 0, len(chars)-1)

	return string(chars)
}

func ReverseWord(word []byte, start, end int) {
	if word == nil {
		return
	}

	for {
		if start >= end {
			break
		}

		temp := word[start]
		word[start] = word[end]
		word[end] = temp
		start++
		end--
	}
}
