package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.Mutex：互斥锁，等同于linux下的pthread_mutex_t
//多个线程同时运行，获得Mutex锁者线程优先执行，其余线程阻塞等待

// sync.Mutex为互斥锁，同一时间只能有一个goroutine获得互斥锁。
// 使用Lock()加锁，Unlock()解锁，加锁前不能解锁，加锁后不能继续加锁。
// 已经锁定的 Mutex 并不与特定的 goroutine 相关联，可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁。
// 适用于同一时间只能有一个goroutine访问资源的场景。
// 下面的代码如果不使用Mutex，输出的会是1 1 2 2 3 3而不是1 2 3 1 2 3。

// 下面的代码如果不使用Mutex，输出的会是1 1 2 2 3 3而不是1 2 3 1 2 3。
var (
	mutex sync.Mutex
)

func print123() {
	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < 3; i++ {
		fmt.Println(i + 1)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go print123()
	go print123()
	time.Sleep(time.Second * 5)
}
