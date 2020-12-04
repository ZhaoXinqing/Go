package main

import (
	"context"
	"fmt"
	"time"
)

//定义一个新类型，包含一个Context 字段
type otherContext struct {
	context.Context
}

func main() {
	//使用context.Background 构建一个withCancel 类型的上下文
	ctxa, work1Cancel := context.WithCancel(context.Background())

	//work模拟退出通知
	go work(ctxa, "work1")

	//使用withDeadline 包装前面的上下文对象ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	//work模拟超时通知
	go work(ctxa, "work2")

	//使用with 对象包装前面的上下文对象ctxb
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "this is some things")

	go workWithValue(ctxc, "work3")

	//休眠10s,让work2 、work3 退出

	time.Sleep(5 * time.Second)

	//显示调用work1 的cancel 方法通知其退出
	work1Cancel()
	fmt.Println("================")
	fmt.Println("work1 exec cancel ...")
	fmt.Println("================")
	//等待work1 打印退出信息
	time.Sleep(2 * time.Second)
	fmt.Println("all things done")
}

//模拟逻辑处理
func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

//等待前端的退出通知，并试图获取context 传递的数据
func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s\n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
