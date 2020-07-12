package main

import (
	"fmt"
	"sort"
)
type People struct {
	Name string
	Age  int
	
type Peoples []People

func main() {

	// 声明一个简单的结构体
	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// 任意切片排序，但是切片必须包含Len(),Swap(),Less()方法
	sort.Sort(peoples)
	fmt.Println(peoples)

	// 检查是否已经被排序
	fmt.Println(sort.IsSorted(peoples))
}
