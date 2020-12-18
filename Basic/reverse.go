package Basic

import (
	"fmt"
	"strings"
)

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// reverseSlice 切片反序
func reverseSlice(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// reverseWords 单词反转
func reverseWords(s string) string {
	es := strings.Split(s, " ")
	es1 := make([]string, 0, len(es))
	for i := len(es) - 1; i >= 0; i-- {
		if 0 != len(es[i]) {
			es1 = append(es1, es[i])
		}
	}
	return strings.Join(es1, " ")
}

func main() {
	testString := "abcdefg"
	testSlice := []byte{11, 22, 33, 44}
	testWords := "haha , hello world !"

	String1 := reverseString(testString)
	fmt.Println(String1)

	Slice1 := reverseSlice(testSlice)
	fmt.Println(Slice1)

	Words1 := reverseWords(testWords)
	fmt.Println(Words1)

}
