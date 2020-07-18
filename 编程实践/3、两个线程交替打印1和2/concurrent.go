// 目的：Golang两个线程实现交替打印1和2 一共10次
// 思路：一个进程写，一个进程读，一个管道存数据，一个管道进行退出控制
package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)

	go func(intChan chan int) {
		for i := 0; i < 20; i++ {
			intChan <- (i%2 + 1)
		}
		close(intChan)
	}(intChan)

	go func(intChan chan int, exitChan chan bool) {
		for {
			v, ok := <-intChan
			if ok {
				fmt.Println(v)w
			} else {
				exitChan <- true
				close(exitChan)
				break
			}
		}
	}(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

// 输出
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
// 1
// 2
