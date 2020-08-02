// Golang 自带的数据结构是没有set集合的。
// 那么，今天我们通过map来实现一个不重复的set集合。

// 根据go中map的keys的无序性和唯一性，可以将其作为set
// golang实现set集合,变相实现切片去重
package main

import "fmt"

func main() {
	hashSet := make(map[string]struct{})
	data := []string{"Hello", "World", "213", "3213", "213", "World"}
	for _, v := range data {
		hashSet[v] = struct{}{}
	}
	for k := range hashSet {
		fmt.Println(k)
	}
}

// output
// Hello
// World
// 213
// 3213
