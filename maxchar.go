package dsalgo

func MaxChar(str string) string {
	maxChar := ""
	charMap := make(map[string]int)
	for _, charRune := range str {
		c := string(charRune)
		if charMap[c] != 0 {
			charMap[c] += 1
		} else {
			charMap[c] = 1
		}
	}

	var max = 0
	for k, v := range charMap {
		if v > max {
			maxChar = k
			max = v
			continue
		}
	}
	return maxChar
}
