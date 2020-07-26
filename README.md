## 简介 ##
Go语言是一门极具活力的新生代语言，它的设计就是为高并发而服务的，并集合了当前诸多优秀程序设计语言的优点，如 **自动垃圾回收、快速的静态类型检查** 和 **高并发性能**等；

## Go 语言中文开源图书、资料或文档:
http://books.studygolang.com/



## Go语言常见的坑
https://studygolang.com/articles/16949?fr=sidebar

## go by example
https://books.studygolang.com/gobyexample/



- 依赖管理：https://www.cnblogs.com/Dr-wei/p/11742253.html
- 内存管理：https://www.cnblogs.com/qingaoaoo/p/13296038.html
- 编码规范：https://github.com/xxjwxc/uber_go_guide_cn

　　
## 手写洗牌


说：Java是企业开发语言，而Go是系统编程语言

- 为什么Go比Java更快

Go被编译为机器代码并直接执行，这使得它比Java快得多。之所以如此，是因为Java使用VM来运行其代码，这使得它与Golang相比变慢。Golang在内存管理方面也很出色，这在编程语言中至关重要。Golang没有引用但有指针。在局部变量的情况下，与Java相比，Golang更好。局部变量在Java语言和其他语言中是存储在堆栈中。但是在Golang中有一个例外，它可以从函数返回指向局部变量的指针