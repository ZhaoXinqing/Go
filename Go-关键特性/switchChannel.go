package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go send(ch2, 10)

	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Printf("get value %d from ch1\n", val)
		case val := <-ch2:
			fmt.Println("get value %d from ch2\n", val)
		case <-time.After(2 * time.Second):
			fmt.Println("Time out")
			return
		}
	}
}
