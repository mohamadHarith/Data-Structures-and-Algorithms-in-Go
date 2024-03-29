package main

import (
	"log"
	"reflect"
)

func main() {
	// log.Println(NaivePatternSearch("THIS IS A TEXT TEXT text TEXT", "TEXT"))

	s := "年test"

	a := s[0]

	log.Println(a)
	log.Println(len(s))
	log.Println(reflect.TypeOf(a))
}

func NaivePatternSearch(text, pattern string) []int {
	var foundIndex, patternIndex int = -1, 0

	var foundIndices []int

	// findPattern:
	// iterate over the text
	for i := 0; i < len(text); i++ {
		// do a while loop on the pattern
		for {
			if patternIndex > len(pattern) {
				patternIndex = 0
				break
			}
			// if there is a match
			if text[i] == pattern[patternIndex] {
				// log.Println("matched text=>", i, string(text[i]))
				// log.Println("matched pattern=>", patternIndex, string(pattern[patternIndex]))

				// set found index
				if foundIndex == -1 {
					foundIndex = i
				}

				// if end of pattern is reached, reset the pattern index
				if patternIndex == len(pattern)-1 {
					foundIndices = append(foundIndices, foundIndex)
					patternIndex = 0
					foundIndex = -1
					break

				} else {
					// else increment the pattern index
					patternIndex++
				}

				// whatever the case, always break the inner loop
				// and go to next char in the text
				break
			} else {
				// if there isnt match
				patternIndex = 0
				foundIndex = -1
				break
			}
		}

	}

	return foundIndices
}
