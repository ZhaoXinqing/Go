package main

import "fmt"

// 递归就是重复调用同一个函数过程来解析相同的问题
func recursion() {
	resursion()
}
func main() {
	resursion()
}

// 使用递归时，一定要注意退出条件，否则就会出现死循环！

// 示例
func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", fibonaci(i))
	}
}

// 递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
// 实现阶乘：
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

// 实现斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

//  fibonacci 实现，用到多返回值特性，降低复杂度：
