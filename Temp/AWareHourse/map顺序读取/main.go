// map如何顺序读取?
// 可以通过sort中的排序包进行对map中的key进行排序
// golang:使用 sort 来排序

package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = map[string]int{
		"d_hello":   0,
		"a_morning": 1,
		"b_my":      2,
		"c_girl":    3,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
