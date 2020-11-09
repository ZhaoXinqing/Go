package main

import "fmt"

func main() {
	a := []int64{10, 20, 10, 20, 30, 40}
	a = RemoveRepByMap(a)
	fmt.Println(a)
}
