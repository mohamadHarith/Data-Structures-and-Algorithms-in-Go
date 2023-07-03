package strings

func NaivePatternSearch(text, pattern string) int {
	var foundIndex, patternIndex int = -1, 0

	// iterate over the text
	for i := 0; i < len(text); i++ {
		// do a while loop on the pattern
		for {
			if patternIndex > len(pattern) {
				break
			}
			// if there is a match
			if text[i] == pattern[patternIndex] {
				// set found index
				if foundIndex == -1 {
					foundIndex = i
				}

				// if end of pattern is reached, reset the pattern index
				if patternIndex == len(pattern)-1 {
					patternIndex = 0
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

		// for j := 0; j < len(pattern); j++ {
		// 	if text[i] == pattern[j] {
		// 		foundIndex = i
		// 		if j == len(pattern)-1 {
		// 			break findPattern
		// 		}
		// 		continue
		// 	} else {
		// 		foundIndex = -1
		// 	}
		// }
	}

	return foundIndex
}
