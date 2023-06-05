package main

import "log"

func main() {
	var a = "hello,  world"

	var chars = []byte(a)

	start := 0
	for end := 0; end < len(a); end++ {
		if a[end] == ' ' {
			ReverseWord(chars, start, end-1)
			start = end + 1
		}
	}

	// reverse the last word
	ReverseWord(chars, start, len(a)-1)

	// reverse the entire string
	ReverseWord(chars, 0, len(a)-1)

	log.Println(string(chars))

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
