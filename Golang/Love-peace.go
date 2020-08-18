// 使用闭包特性实现一个计数器
func createCounter(initial int) func() int {
	if initial < 0{
		initial = 0
	}
	return func() int {
		initial ++
		return initial;
	}
}

// func main()
cl := createCounter(1)
fmt.Println(c1())  // 2
fmt.Println(c1()) // 3

c2 := createCounter(10)
fmt.Println(c2()) // 11
fmt.Println(cq())  // 4

