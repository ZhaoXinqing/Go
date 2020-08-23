package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

// GenerateInt 扇入
func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			case ch <- <-GenerateIntA():
			case ch <- <-GenerateIntB():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateInt()

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
