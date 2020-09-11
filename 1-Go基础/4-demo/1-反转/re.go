package rev

import "strings"

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// 切片反转
func reverseSlice(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 单词反转
func reverseWords(s string) string {
	es := string.Split(s, " ")
	es1 := make([]string, 0, len(es))
	for i := len(es) - 1; i >= 0; i-- {
		if 0 != len(es[i]) {
			es1 = append(es1, es[i])
		}
	}
	return strings.Join(es1, " ")
}
