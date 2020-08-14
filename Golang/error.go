// golang的所有错误都是通过 error 来处理的：

// 一，经典的处理方式：
err := xxxx
if err != nil {
	// ...
}

// 二，创建自己的 error信息，把 err 封装成各种类型：
func Sqrt(value float64)(float64,error){
	if (value < 0) {
		return 0, error.New("Math:negative number passed to Sqrt")
	}
	return math.Sqrt(value)
}

// Go语言通过内置的错误接口提供了非常简洁的错误处理机制
// err是一个接口类型，它的定义为：
type error interface {
	Error() string
}
// 我们可以在编码中通过error接口类型来生成错误信息。
// 函数通常在最后的返回值中返回错误信息。
// 使用errors.New 可返回一个错误信息，示例：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
}


// panic-recover-defer
// panic 用于主动抛出错误，recover 用来捕获 pinic 抛出的错误；

// panic
// 引发 panic：（1）程序主动调用，（2）程序产生运行时错误，由运行时检测出并退出；
// 发生 panic后，程序从调用或者发生panic的地方立即返回，逐层向上执行函数的defer语句，
// 		- 直到被recover 捕获或运行到函数最外层；
// panic 在没有用 recover 前以及在 recover 捕获那一级函数栈，panic 之后的代码均不会执行；
// 		- 一旦被 recover 捕获后，外层的函数栈代码恢复正常，所有代码均会得到执行；
// 利用 recover 捕获 panic 时，defer 需要再 panic 之前声明，
// 		- 否则由于 panic 之后的代码得不到执行，因此也无法 recover；
// panic 后，不再执行后面的代码，立即按照逆序执行 defer，并逐级往外层函数栈扩散；
// 		- defer 就类似 finally；

// recover用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，
// defer只有在后面的函数体内直接被调用才能捕获panic终止异常，否则返回nil，异常继续向外传递。



