// 协程使用channel 接受和发送数据。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	ch := make(chan string)
	go printInput(ch)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("End the game!")
			break
		}
	}
	defer close(ch)
}
