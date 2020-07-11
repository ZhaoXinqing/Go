//只执行一次以后不再触发

// sync.Once 是 Golang package 中使方法只执行一次的对象实现，作用与** init **函数类似。但也有所不同。

// init 函数是在文件包首次被加载的时候执行，且只执行一次
// sync.Onc 是在代码运行中需要的时候执行，且只执行一次
// 当一个函数不希望程序在一开始的时候就被执行的时候，我们可以使用 sync.Once。

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// 输出：
// # Output:
// Only once
// sync.Once 使用变量 done 来记录函数的执行状态，使用 sync.Mutex 和 sync.atomic 来保证线程安全的读取 done 。
