package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("receive:", v)
	}
}

func main() {
	queue := make(chan int, 1)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(1e9) //让Producer与Consumer完成
}

// 1. 解耦 （ 降低生产者 和 消费者之间 耦合度 ）
// 2. 并发 （生产者消费者数量不对等时，能保持正常通信）
// 3. 缓存 （生产者和消费者 数据处理速度不一致时， 暂存数据）
