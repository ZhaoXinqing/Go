package main

import (
	"fmt"
	"runtime"
)

var ch chan string = make(chan string)

var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
var ch3 chan int = make(chan int)
var ch4 chan int = make(chan int)

func sendA() {
	for {
		c := <-ch1
		if c == 1 {
			ch <- "A"
			ch2 <- 1
		}
		runtime.Gosched()
	}
}

func sendB() {
	for {
		c := <-ch2
		if c == 1 {
			ch <- "B"
			ch3 <- 1
		}
		runtime.Gosched()
	}

}

func sendC() {
	for {
		c := <-ch3
		if c == 1 {
			ch <- "C"
			ch4 <- 1
		}
		runtime.Gosched()
	}
}

func sendD() {
	for {
		c := <-ch4
		if c == 1 {
			ch <- "D"
			ch1 <- 1
		}
		runtime.Gosched()
	}
}

func main() {
	go sendA()
	go sendB()
	go sendC()
	go sendD()
	ch1 <- 1

	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		if n == 0 {
			continue
		}

		for i := 0; i < n*4; i++ {
			str := <-ch
			fmt.Printf("%s", str)
		}
		fmt.Println()
	}
}
