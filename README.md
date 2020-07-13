## 简介 ##

Go是一种新的语言，一种并发的、带垃圾回收的、快速编译的语言。它具有以下特点：

- [内置并发编程支持](https://www.jianshu.com/p/63dbec263d2a)
- 使用协程（goroutine）做为基本的计算单元。轻松地创建协程。
- 使用数据通道（channels）来实现协程间的同步和通信。
- 内置了映射（map）和切片（slice）类型。
- [支持多态（polymorphism）](https://blog.csdn.net/jw915086731/article/details/86751334)。
- 使用接口（interface）来实现装盒（value boxing）和反射（reflection）。
- [支持指针。](http://c.biancheng.net/view/21.html)
- 支持函数闭包（closure）。
- 支持方法。
- 支持延迟函数调用（defer）。
- 支持类型内嵌（type embedding）。
- 支持类型推断（type deduction or type inference）。
- [内存安全。](https://blog.csdn.net/wenrennaoda/article/details/95935355)
- 自动垃圾回收。
- 良好的代码跨平台性。

Go是一种编译型静态语言，它结合了解释型语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。目的成为现代的，支持网络与多核计算的语言。具有轻量级的类型系统，并发与垃圾回收机制，严格的依赖规范等等。

可读性是在Go语言的设计中一个非常重要的考虑因素。 一个Go程序员常常可以轻松读懂其他Go程序员写的代码。 甚至对于一个没有Go编程经验但具有其它语言编程经验的程序员来说，读懂一份Go源码也不是一件难事。

目前，使用最广泛的Go编译器由Go官方设计和开发团队维护的标准编译器。标准编译器也常常称为gc（是Go compiler的缩写，不是垃圾回收garbage collection的缩写）。 Go官方设计和开发团队也维护着另外一个编译器，gccgo。 gccgo是gcc编译器项目的一个子项目。 gccgo的使用广泛度大不如gc， 它的主要作用是做为一个参考，来保证gc的实现正确性。 目前两个编译器的开发都很活跃，尽管Go开发团队在gc的开发上花费的精力更多。

自从Go语言正式发布后，Go的语法变化很小。 但是标准编译器gc却在不断地提升改进。 


**go的内建接口只有error**

