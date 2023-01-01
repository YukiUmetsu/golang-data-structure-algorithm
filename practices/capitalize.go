package dsalgo

import "strings"

func Capitalize(s string) string {
	strArr := strings.Split(s, "")
	for i := range strArr {
		if i+1 == len(strArr) {
			break
		}
		if i == 0 || strArr[i-1] == " " {
			strArr[i] = strings.ToUpper(strArr[i])
		}
	}
	return strings.Join(strArr, "")
}
