// Go内建的map不是线程安全的，所以之前都会使用加锁的方式控制对map的并发访问，
// 而在新版本中可以使用sync.Map代替map+RWMutex，sync.Map相比后者有更好的性能，
// 下面是sync.Map的基本使用方法：

package main

import (
	"fmt"
	"sync"
)

var (
	players sync.Map
)

func main() {
	//设置key对应的value
	players.Store("xiao ming", 100)
	//返回key对应的valuee，value不存在时返回nil,false
	if value, ok := players.Load("xiao ming"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("key:xiao ming does not exist!")
	}
	//如果key对应的value存在，则返回value和true
	//如果key对应的value不存在，则将key对应的value设置为参数中的value，并返回value和false
	if value, ok := players.LoadOrStore("xiao hong", 120); ok {
		fmt.Println(value)
	} else {
		fmt.Println("key:xiao hong does not exist，store it!")
	}
	//遍历map，函数返回false时停止遍历
	players.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//删除key对应的value
	players.Delete("xiao hong")
}
