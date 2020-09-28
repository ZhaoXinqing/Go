## 1、使用信道来标记完成

信道可以实现多个协程间的通信，那么我们只要定义一个信道，在任务完成后，往信道中写入true，然后在主协程中获取到true，就认为子协程已经执行完毕。

```go
import "fmt"
func main() {
    done := make(chan bool)
    go func() {
        for i := 0;i<5;i++ {
            fmt.Println(i)
        }
        done <- true
    }()
    <- done
}
```

## 2、使用WaitGroup

用信道的方法，在单个协程或者协程数少的时候，并不会有什么问题，但在协程数多的时候，代码就会显得非常复杂，更加优雅的方式，使用sync包 提供的 WaitGroup 类型。

WaitGroup  你只要实例化了就能使用，`var 实例名 sync.WaitGroup` ，它的几个方法：

- `Add`：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
- `Done`：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
- `Wait`：阻塞当前协程，直到实例里的计数器归零。

```go
import (
	"fmt"
    "sync"
)
unbtqyabiwojhgee
func worker(x int,wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0;i < 5;i++ {
        fmt.Printf("worker %d: %d\n",x,i)
    }
}

func main() {
    var wg sync.WaitGroup
    
    wg.Add(2)
    go worker(1,&wg)
    go worker(2,&wg)
    
    wg.Wait()
}
```


## 3、Context

Context，也叫上下文，它的接口定义如下

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```


```go
package main

import (
	"context"
    "fmt"
    "time"
)

func monitor(ctx context.Context, number int) {
    for {
        select {
        case v := <- ctx.Done():
            fmt.Printf("监控器%v,正在监控中。。。\n",number,v)
            return
        default:
            fmt.Printf("监控器%v,正在监控中...\n", number)
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx,cancel := context.WithCancel(context.Background())
    
    for i := 1; i <= 5; i++ {
        go monitor(ctx,i)
    }
    time.Sleep( 1 * time.Second)
    // 关闭所有 goroutine
    cancel()
    
    // 等待5s，若此时屏幕没有输出<正在监控> 就说明说有的goroutine都已经关闭了
    time.Sleep( 5 * time.Second)
    fmt.Println("主程序退出！！")
}
```

可以看到 Context 接口共有 4 个方法

- `Deadline`：返回的第一个值是 **截止时间**，到了这个时间点，Context 会自动触发 Cancel 动作。返回的第二个值是 一个布尔值，true 表示设置了截止时间，false 表示没有设置截止时间，如果没有设置截止时间，就要手动调用 cancel 函数取消 Context。
- `Done`：返回一个只读的通道（只有在被cancel后才会返回），类型为 `struct{}`。当这个通道可读时，意味着parent context已经发起了取消请求，根据这个信号，开发者就可以做一些清理动作，退出goroutine。
- `Err`：返回 context 被 cancel 的原因。
- `Value`：返回被绑定到 Context 的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。