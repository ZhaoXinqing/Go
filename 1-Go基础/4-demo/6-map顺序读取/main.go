package main

import (
	"fmt"
	"sort"
)

func main() {
	/* 声明索引类型为字符串的map */
	var testMap = make(map[string]string)
	testMap["Bda"] = "B"
	testMap["Ada"] = "A"
	testMap["Dda"] = "D"
	testMap["Cda"] = "C"
	testMap["Eda"] = "E"

	var testSlice []string
	for key := range testMap {
		testSlice = append(testSlice, key)
	}

	/* 对slice数组进行排序，然后就可以根据key值顺序读取map */
	sort.Strings(testSlice)
	fmt.Println("排序输出:")
	for _, Key := range testSlice {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := testMap[Key]; ok {
			fmt.Println(Key, ":", Value)
		}
	}
}
