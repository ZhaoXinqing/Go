package main

import (
	"fmt"
)

// Person ...
type Person struct {
	//结构也是一中类型
	Name string //定义struct的属性
	Age  int
}

func main() {
	a := Person{}
	a.Name = "joe" //对struct的属性进行操作，类型与class的使用方法
	a.Age = 19
	fmt.Println(a)

	var n map[int64]interface{}
	n = make(map[int64]interface{})
	n[1] = a
	fmt.Println(n)
}
