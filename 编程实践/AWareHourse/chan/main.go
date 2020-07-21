package main

import "fmt"

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
