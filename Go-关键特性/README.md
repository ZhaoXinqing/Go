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
是在计算机科学中，类型系统用于定义如何将编程语言中的数值和表达式归类为许多不同的类型，如何操作这些类型，这些类型如何互相作用。类型可以确认一个值或者一组值具有特定的意义和目的（虽然某些类型，如抽象类型和函数类型，在程序运行中，可能不表示为值）。类型系统在各种语言之间有非常大的不同，也许，最主要的差异存在于编译时期的语法，以及运行时期的操作实现方式。

更多了解：https://baike.baidu.com/item/%E7%B1%BB%E5%9E%8B%E7%B3%BB%E7%BB%9F/4273825?fr=aladdin

## golang-map是线程安全的吗？
https://studygolang.com/articles/23184?fr=sidebar

## 深度解密Go语言之map ##
https://www.cnblogs.com/qcrao-2018/p/10903807.html


## go的值传递和引用
值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数


**<p align = right> --凡事有因才有果，任何事情最终还是要靠自己来解决。**</p>