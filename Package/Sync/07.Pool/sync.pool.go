package main

import (
	"fmt"
	"sync"
)

// sync.Pool 对象缓存
// 		对象获取
// 			尝试从私有对象获取(私有对象是协程安全的)
// 			私有对象不存在 尝试从当前Processor的共享池获取(共享池是协程不安全 需要锁)
// 			如果当前Processor共享池是空的 那么尝试去其他Processor的共享池获取
// 			如果所有子池都是空的 最后使用用户指定的New()函数产生一个新的对象返回
// 		对象放回
// 			如果私有对象不存在则保存为私有对象
// 			如果私有对象存在 放入当前Processor子池的共享池中
// 		对象的声明周期
// 			GC会清除sync.Pool缓存的对象
// 			对象的缓存有效期为下一次GC之前
func main() {
	var wg sync.WaitGroup
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new obj.")
			return 100
		},
	}
	pool.Put(10)
	pool.Put(10)
	pool.Put(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer func() {
				wg.Done()
			}()
			fmt.Println(pool.Get())
		}(i)
	}
	wg.Wait()
}
