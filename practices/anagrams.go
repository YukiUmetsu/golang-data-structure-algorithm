package dsalgo

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func IsAnagram(str1, str2 string) bool {
	var err error
	str1, err = filterString(str1)
	if err != nil {
		fmt.Printf("failed to filter str1 %v", err)
		return false
	}

	str2, err = filterString(str2)
	if err != nil {
		fmt.Printf("failed to filter str2 %v", err)
		return false
	}

	charMap1 := createCharMap(str1)
	charMap2 := createCharMap(str2)

	// deep compare two maps with reflection
	return reflect.DeepEqual(charMap1, charMap2)
}

func createCharMap(str string) map[string]int {
	charMap := make(map[string]int)
	for _, char := range str {
		letter := string(char)
		if _, ok := charMap[letter]; ok {
			charMap[letter] += 1
		} else {
			charMap[letter] = 1
		}
	}
	return charMap
}

func filterString(str string) (string, error) {
	// everything except those characters [0-9A-Za-z_!]
	re, err := regexp.Compile(`[^0-9A-Za-z_!]*`)
	if err != nil {
		fmt.Printf("failed to compile regex err: %v\n", err)
		return str, err
	}
	res := re.ReplaceAllString(str, "")
	return strings.ToLower(res), nil
}
