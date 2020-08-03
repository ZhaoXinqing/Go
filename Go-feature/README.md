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

## 匿名函数和闭包
Go 支持通过 闭包来使用 匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包
https://blog.csdn.net/ycy258325/article/details/54632915

## 并发
独特的MPG的线程模型和CSP并发理念也为Go语言高性能、高并发的特性增分不少；
https://www.cnblogs.com/qingaoaoo/p/13295835.html

## Gomodule
Go module的出现有效解决了Go语言依赖混乱的问题；

## 反射
反射机制就是在程序运行时动态调用对象的方法和属性，标准库reflect提供了相关的功能。在reflect包中，通过reflect.TypeOf()，reflect.ValueOf()分别从类型、值的角度来描述一个Go对象。
实际上，反射是通过检查一个接口的值，变量首先被转换成空接口

## 接口
接口 是 方法特征 的命名集合
在Go语言的实现中，一个interface类型的变量存储了2个信息, 一个《值，类型》对(value, type)
value是实际变量值，type是实际变量的类型。

## golang-map
https://studygolang.com/articles/23184?fr=sidebar


## go的值传递和引用
值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数

Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。


## 强类型的静态编译型语言
https://blog.csdn.net/ling12abc/article/details/102993484


**<p align = right> --靠自己**</p>