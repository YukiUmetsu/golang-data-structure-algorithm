package dsalgo

import "strings"

func NumberOfVowels(str string) int {
	vowels := "aeiouAEIOU"
	counter := 0
	for _, strRune := range str {
		if strings.Contains(vowels, string(strRune)) {
			counter++
		}
	}
	return counter
}
