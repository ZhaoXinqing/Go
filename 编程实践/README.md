## map如何顺序读取? ##
可以通过sort中的排序包进行对map中的key进行排序
golang:使用 sort 来排序

	package main
	import (
	    "fmt"
	    "sort"
	)
	func main() {
	    var m = map[string]int{
	        "hello":         0,
	        "morning":       1,
	        "my":            2,
	        "girl":   		3,
	    }
	    var keys []string
	    for k := range m {
	        keys = append(keys, k)
	    }
	    sort.Strings(keys)
	    for _, k := range keys {
	        fmt.Println("Key:", k, "Value:", m[k])
	    }
	}

## map实现set ##
	// Golang 自带的数据结构是没有set集合的。
	// 那么，今天我们通过map来实现一个不重复的set集合。
	
	// 根据go中map的keys的无序性和唯一性，可以将其作为set
	// golang实现set集合,变相实现切片去重
	package main
	
	import "fmt"
	
	func main() {
		hashSet := make(map[string]struct{})
		data := []string{"Hello", "World", "213", "3213", "213", "World"}
		for _, v := range data {
			hashSet[v] = struct{}{}
		}
		for k := range hashSet {
			fmt.Println(k)
		}
	}
	
	// output
	// Hello
	// World
	// 213
	// 3213


## 点赞系统设计
基本的设计思路有大致两种， 一种自然是用mysql等
数据库直接落地存储， 另外一种就是利用点赞的业务特征来扔到redis(或memcache)中， 然后离线刷回mysql等。
https://blog.csdn.net/weixin_34406086/article/details/86117916
https://blog.csdn.net/qq_34264849/article/details/84401198

## 一个大整数（字符串形式表示的），移动字符求比它大的数中最小的

## 所有左叶子节点的和

## 手写循环队列
写的循环队列是不是线程安全，不是，怎么保证线程安全，加锁，效率有点低啊，然后面试官就提醒Go推崇原子操作和channel

m个n大小的有序数组求并集，一开始是2路归并，求时间复杂度，后来在面试官提醒直接m路归并，求时间复杂度


