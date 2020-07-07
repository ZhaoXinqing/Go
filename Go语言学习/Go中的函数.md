## 1. 关于函数 ##
函数是基于功能或 逻辑进行封装的可复用的代码结构。将一段功能复杂、很长的一段代码封装成多个代码片段（即函数），有助于提高代码可读性和可维护性。

在 Go 语言中，函数可以分为两种：

- 带有名字的普通函数
- 没有名字的匿名函数

由于 Go语言是编译型语言，所以函数编写的顺序是无关紧要的，它不像 Python 那样，函数在位置上需要定义在调用之前。


## 2. 函数的声明 ##
函数的声明，使用 func 关键字，后面依次接 函数名，参数列表，返回值列表，用 {} 包裹的代码逻辑体

	func 函数名(形式参数列表)(返回值列表){
	    函数体
	}

- 形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由参数调用者提供
- 返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。

举个例子，定义一个 sum 函数，接收两个 int 类型的参数，在运行中，将其值分别赋值给 a，b，并规定必须返回一个int类型的值 。

	func sum(a int, b int) (int){
	    return a + b}func main() {
	fmt.Println(sum(1,2))}

## 3. 函数实现可变参数 ##


上面举的例子，参数个数都是固定的，这很好理解 ，指定什么类型的参数就传入什么类型的变量，数量上，不能多一个，也不能少一个。（好像没有可选参数）。
在 Python 中我们可以使用 args 和 *kw ，还实现可变参数的函数。

可变参数分为几种：

- 多个类型一致的参数
- 多个类型不一致的参数

**1、多个类型一致的参数**

首先是多个类型一致的参数。
这边定义一个可以对多个数值进行求和的函数，
使用 ...int，表示一个元素为int类型的切片，用来接收调用者传入的参数。

// 使用 ...类型，表示一个元素为int类型的切片func sum(args ...int) int {
    var sum int
    for _, v := range args {
        sum += v
    }
    return sum}func main() {
    fmt.Println(sum(1, 2, 3))}
// output: 6

其中 ... 是 Go 语言为了方便程序员。/代码而实现的语法糖，如果该函数下会多个类型的函数，这个语法糖必须得是最后一个参数。
同时这个语法糖，只能在定义函数时使用。
2、多个类型不一致的参数
上面那个例子中，我们的参数类型都是 int，如果你希望传多个参数且这些参数的类型都不一样，可以指定类型为 ...interface{}，然后再遍历。
比如下面这段代码，是Go语言标准库中 fmt.Printf() 的函数原型：

	import "fmt"func MyPrintf(args ...interface{}) {
	    for _, arg := range args {
	        switch arg.(type) {
	            case int:
	                fmt.Println(arg, "is an int value.")
	            case string:
	                fmt.Println(arg, "is a string value.")
	            case int64:
	                fmt.Println(arg, "is an int64 value.")
	            default:
	                fmt.Println(arg, "is an unknown type.")
	        }
	    }}
	func main() {
	    var v1 int = 1
	    var v2 int64 = 234
	    var v3 string = "hello"
	    var v4 float32 = 1.234
	    MyPrintf(v1, v2, v3, v4)}

在某些情况下，我们需要定义一个参数个数可变的函数，具体传入几个参数，由调用者自己决定，但不管传入几个参数，函数都能够处理。
比如这边实现一个累加

	func myfunc(args ...int) {
	    for _, arg := range args {
	        fmt.Println(arg)
	}}

## 4. 多个可变参数函数传递参数 ##

上面提到了可以使用 ... 来接收多个参数，除此之外，它还有一个用法，就是用来解序列，将函数的可变参数（一个切片）一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身。
同样这个用法，也只能在给函数传递参数里使用。

例子如下：

	import "fmt"
	func sum(args ...int) int {
	    var result int
	    for _, v := range args {
	        result += v
	    }
	    return result}
	func Sum(args ...int) int {
	    // 利用 ... 来解序列
	    result := sum(args...)
	    return result}func main() {
	    fmt.Println(sum(1, 2, 3))}
