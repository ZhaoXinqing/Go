// 通过两个 goroutine 来计算数字之和，

package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{3, 5, 6, -4, 89, 10}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
