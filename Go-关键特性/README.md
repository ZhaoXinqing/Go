Go是一种新的语言，一种并发的、带垃圾回收的、快速编译的语言。它具有以下特点：
- [内置并发编程支持](https://www.jianshu.com/p/63dbec263d2a)
- 使用协程（goroutine）做为基本的计算单元。轻松地创建协程。
- 使用数据通道（channels）来实现协程间的同步和通信。
- 内置了映射（map）和切片（slice）类型。
- [支持多态（polymorphism）](https://blog.csdn.net/jw915086731/article/details/86751334)。
- 使用接口（interface）来实现装盒（value boxing）和反射（reflection）。
- [支持指针。](http://c.biancheng.net/view/21.html)
- 支持函数闭包（closure）。https://www.cnblogs.com/hzhuxin/p/9199332.html
- 支持方法。
- 支持延迟函数调用（defer）。
- 支持类型内嵌（type embedding）。
- 支持类型推断（type deduction or type inference）。
- [内存安全。](https://blog.csdn.net/wenrennaoda/article/details/95935355)
- 自动垃圾回收。
- 良好的代码跨平台性。

## 并发
独特的MPG的线程模型和CSP并发理念也为Go语言高性能、高并发的特性增分不少；
https://www.cnblogs.com/qingaoaoo/p/13295835.html

## Gomodule
Go module的出现有效解决了Go语言依赖混乱的问题；

## 反射
提供给开发商在运行时对代码进行修改的能力，方便各种框架（如注入依赖框架）的实现；

## 单元测试

## 内存管理
（垃圾回收、逃逸分析、内存泄漏）

## 类型系统
https://baike.baidu.com/item/%E7%B1%BB%E5%9E%8B%E7%B3%BB%E7%BB%9F/4273825?fr=aladdin

## golang-map是线程安全的吗？
https://studygolang.com/articles/23184?fr=sidebar

## 深度解密Go语言之map ##
https://www.cnblogs.com/qcrao-2018/p/10903807.html


## go的值传递和引用
值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
**默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数**

## 方法
Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。

## 接口
接口 是 方法特征 的命名集合

## 错误处理
Go 语言使用一个独立的·明确的返回值来传递错误信息的。这与使用异常的 Java 和 Ruby 以及在 C 语言中经常见到的超重的单返回值/错误值相比，Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用。

按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
## 强类型的静态编译型语言
https://blog.csdn.net/ling12abc/article/details/102993484

**<p align = right> --凡事有因才有果，任何事情最终还是要靠自己来解决。**</p>

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


## 一、变量类型 ##
变量分为值类型，指针类型和引用类型。

- 在golang中故意淡化了指针的概念，我们只需要关注值类型和引用类型就可以。你在官方介绍中也很少看到指针类型这一概念

## 二、值类型和引用类型的区别 ##
- 值类型变量：除开slice,map, channel类型之外的变量都是值类型
- 引用类型变量：slice, map, channel这三种。
### 1，零值不同 ###
- 指针类型的变量，零值都是nil。
- 值类型的变量，零值是其所在类型的零值。
	1. int32类型的零值是0
	1. string类型的零值是""
	1. bool类型的零值是false
	1. 符合结构struct类型的零值是其每个成员的零值的组合

### 2，变量申明后是否需要初始化才能使用
- 指针类型的变量，需要初始化才能使用。(slice是一个特例，slice的零值是nil，但是可以直接append)
- 值类型的变量，不用初始化，可以直接使用（struct声明后可以直接使用直接赋值，但是map就不行）
###3，初始化方法不同

- 值类型的变量，其实不需要初始化就可以使用。如果有良好的代码习惯，使用前进行初始化也是非常提倡的。
	- 基本类型的初始化非常简单:
		- var i int; i = 1;
		- var b bool; b = true
		- var s string; s = ""
	- 符合类型struct的初始化有两种:
		- s1 = Student1{}
		- s1 = new(Student1)


- 引用类型的变量，初始化方式也不一样：
	- slice类型，用make，new，{}都可以
	- map类型，用make，new，{}都可以
	- channel类型，只能用make活着new初始化

	

	//map可以用{}，make，new三种方式初始化

	- s2 := Student2{} //map[]
	- s2 := make(Student2) //map[]
	- s2 := new(Student2) //&map[]

	//slice可以用{},make,new,但是make的时候需要带len参数
	- type S3 []string
	- s3 := S3{} //[]
	- s3 := new(S3) //&[]
	- s3 := make(S3, 10) //[ ]

	//channel只能用make或者new

	- type Student4 chan string
	- s4 := new(S4) //0xc000096000
	- s4 := make(S4) //0xc000082008
	- s4 := S4{} //编译器报错：invalid type for composite literal: Student4

## 三、make和new的区别


- make返回的是对象。
	- 对值类型对象的更改，不会影响原始对象的值
	- 对引用类型对象的更改，会影响原始对象的值
- new返回的是对象的指针，对指针所在对象的更改，会影响指针指向的原始对象的值。

## 四、golang没有引用传递，都是值传递

- 如果函数形参是值类型，则会对值类型做一份拷贝作为函数形参。在函数内对形参变量做的修改，不会影响函数外的那个被传入的变量。
- 如果函数形参是引用类型，则会对引用类型变量做一次拷贝。但是拷贝得到的引用类型变量的值，和被传入调用函数的原始引用类型变量的值，是一样的，即指向的是同一个变量的地址(参考前面值类型变量和引用类型变量图)。所以在函数里面的修改，会影响原始引用变量指向的变量的值。

   type Student struct{
    Age int
	}
   
   
	func main(){
    s := Student{30}
    modify1(s)
    	fmt.Println(s) //{30}
    	
    modify2(&s)
    fmt.Println(s) //{32}
     }
    
     func modify1(s Student){
    s.Age = 31
     }
    
     func modify2(s *Student){
    s.Age = 32
     }

## 什么是内存泄露 ##
内存泄露指的是程序运行过程中已不再使用的内存，没有被释放掉，导致这些内存无法被使用，直到程序结束这些内存才被释放的问题。

Go虽然有GC来回收不再使用的堆内存，减轻了开发人员对内存的管理负担，但这并不意味着Go程序不再有内存泄露问题。在Go程序中，如果没有Go语言的编程思维，也不遵守良好的编程实践，就可能埋下隐患，造成内存泄露问题。

## 怎么发现内存泄露 ##
在Go中发现内存泄露有2种方法，一个是通用的监控工具，另一个是go pprof：


1. 监控工具：固定周期对进程的内存占用情况进行采样，数据可视化后，根据内存占用走势（持续上升），很容易发现是否发生内存泄露。
2. go pprof：适合没有监控工具的情况，使用Go提供的pprof工具判断是否发生内存泄露。

## 方法引用 ##
1：同一个包下公有方法（方法名首字母大写）和可直接引用 <br>2：不通包下的公有方法要加上包名以进行引用

## Goalng方法是依托于结构体的 ##

结构体和其他面向对象语言中的class同等地位。
Golang中方法就是一个包含了接收者的函数，接收者可以是一个命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

语法格式如下：

    func (variable_name variable_data_type) function_name() [return_type]{
    	/* 函数体*/
    }


    type Circle struct {
    	radius float64
    }
    
    func main() {
    	var c Circle
    	c.radius = 10
    	fmt.Println("圆的面积 = ", c.getArea())
    }
    
    func (c Circle) getArea() float64 {
    	return 3.14 * c.radius * c.radius
    }
    
Go 支持通过 闭包来使用 匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包


1、GO语言的匿名函数就是闭包

基本概念：
　　闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。
闭包的价值：
　　闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。

Go语言中的闭包

　　Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。例：

// closure.go
package main

import (
    "fmt"
)

func main() {
    var j int = 5

    a := func() func() {
        var i int = 10
        return func() {
            fmt.Printf("i, j: %d, %d\n", i, j)
        }
    }()　　//末尾的括号表明匿名函数被调用，并将返回的函数指针赋给变量a

    a()

    j *= 2

    a()
}
输出：

i, j: 10, 5
i, j: 10, 10

2、Go语言支持匿名函数，即函数可以像普通变量一样被传递或使用。

闭包是“函数”和“引用环境”组成的整体。

func anonymous(n int) func() {
    return func() {
        n++ //对外部变量加1
        fmt.Println(n)
    }
}

func anonymous2(n int) func() {
    sum := n
    a := func() { //把匿名函数作为值赋给变量a(Go不允许函数嵌套，但你可以利用匿名函数实现函数嵌套)
        fmt.Println(sum + 1) //调用本函数外的变量，这里没有()匿名函数不会马上执行
    }
    return a
}

　　fmt.Println("---------anonymous func--------")
    anony := anonymous(10)
    anony()
    anony1 := anonymous(20)
    anony1()
    /**
    *再次调用anony()函数，结果是12，由此得出以下两点
    * 1、内函数对外函数的变量的修改，是对变量的引用
    * 2、变量被引用后，它所在的函数结束，该变量不会马上销毁
    **/
    anony()
    anony1()
    fmt.Println("---------anonymous2----------")
    an := anonymous2(10)
    an()
    an2 := anonymous2(20)
    an2()
    an()
    an2()
输出：

---------anonymous func--------
11
21
12
22
---------anonymous2----------
11
21
11
21

闭包函数出现的条件：
1.被嵌套的函数引用到非本函数的外部变量，而且这外部变量不是“全局变量”;
2.嵌套的函数被独立了出来(被父函数返回或赋值 变成了独立的个体)，  而被引用的变量所在的父函数已结束。
对象是附有行为的数据，而闭包是附有数据的行为。

3、匿名函数执行

　　匿名函数最后括号中加参数，匿名函数立即执行，没有括号或括号中为空，则需要传参。例：

fmt.Println("-----------匿名函数的执行----------")
    m, n := func(i, j int) (m, n int) { // x y 为函数返回值
        return j, i
    }(1, 9) // 直接创建匿名函数并执行

    fmt.Println(m, n)

    f := func(i, j int) (result int) { // f 为函数地址
        result = i + j
        return result
    }

    fmt.Println(f(1, 2))
输出：

-----------匿名函数的执行----------
9 1
3
4、错误： expression in go/defer must be function call

　　go关键字后跟匿名函数报错： expression in go/defer must be function call，在匿名函数后加括号则可。例：

func smtp() error {
    return nil
}
　　go func() {
        err := smtp()
        if err != nil {
            fmt.Printf(err.Error())
        }
    }()
　　若去掉上面代码中的空括号会报错。

　　加()封装成表达式。

　　
