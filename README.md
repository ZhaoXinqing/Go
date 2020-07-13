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

1、Go是一种编译型静态语言，它结合了解释型语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。目的成为现代的，支持网络与多核计算的语言。具有轻量级的类型系统，并发与垃圾回收机制，严格的依赖规范等等。

2、可读性是在Go语言的设计中一个非常重要的考虑因素。 一个Go程序员常常可以轻松读懂其他Go程序员写的代码。 甚至对于一个没有Go编程经验但具有其它语言编程经验的程序员来说，读懂一份Go源码也不是一件难事。

3、自从Go语言正式发布后，Go的语法变化很小。 但是标准编译器gc却在不断地提升改进。 gc（是Go compiler的缩写，不是垃圾回收garbage collection的缩写）。 


**go的内建接口只有error**


## 认识Go语言
由Google公司开发的语言，它兼具编译型语言C、C++等很快的执行速度和翻译型语言Ruby、Python等快速开发的特点，不仅提供了高性能还可以让开发速度更快。
高性能、快速开发：
1、语法简洁，关键字很少方便记忆；
2、类型系统简单高效；
3、编译速度很快，编译的时间非常短；
4、内置并发机制，所以不用被迫使用特定的线程库；
5、自带垃圾回收器，不需要自己管理内存。
快速开发：go语言使用了更加智能的编译器，简化了解决依赖的算法，编译go程序时，不会像Java、C或C++那样遍历依赖链中所有的依赖库。在非常短的时间内就可以完成编译。

（关于游戏开发的Go语言）有人说：
1.go除了后端，其他场景都不合适
2.golang适合写大型并发游戏后端
3.go的出现是为了解决web服务端的并发问题，处理i/o密集，go的协程很香，做游戏服务器还行。写游戏端，大多是计算密集，图形处理，go似乎不太适合。

语言层面定义源代码的格式化
我职业生涯中一些最激烈的辩论发生在团队代码格式的定义上。 Go 通过为代码定义规范格式来解决这个问题。 gofmt 工具会重新格式化您的代码，并且没有选项。不管你喜欢与否，gofmt 定义了如何对代码进行格式化，一次性解决了这个问题。
总结：
Go语言是一门极具活力的新生代语言，它的设计就是为高并发而服务的，并集合了当前诸多优秀程序设计语言的优点，如自动垃圾回收、快速的静态类型检查和高并发性能等；
Go语言是一门静态强类型语言，在程序编译过程中会把变量的反射信息如字段类型、类型信息等写入可执行的过程中。在程序执行的过程中，Go语言虚拟机会加载可执行文件中变量的反射信息，并提供接口用于在运行时获取和修改代码的能力。
工作目录是Go语言项目的开发空间，它是一个目录结构，一般由三个子目录组成
src：包含组成包的源代码，一个目录就是一个包；
pkg：包含编译后的类库；
bin：包含编译后的可执行程序；
而GOPATH是Go语言中使用的一个环境变量，它使用绝对路径提供Go语言项目的工作目录。GOPATH适合处理大量Go语言源码、多个包组合而成的复杂工程。一般一个项目使用一个GOPATH，在编译过程中就不会编译错误的代码和版本；

8、slice，len，cap，共享，扩容
    append函数，因为slice底层数据结构是，由数组、len、cap组成，所以，在使用append扩容时，会查看数组后面有没有连续内存快，有就在后面添加，没有就重新生成一个大的素组

1、包管理：GOPATH、Go Modules
Go语言的代码复用很大程度上依赖于包管理，而包管理很大程度上依赖于环境变量GOPATH。
GOPATH是Go语言中使用的一个环境变量，它使用绝对路径提供Go语言项目的工作目录。它适合处理大量Go语言源码、多个包组合而成的复杂工程。一般建议一个项目使用一个GOPATH，在编译过程中就不会编译错误的代码或者版本；
Go Modules于1.11版本初步引入，在1.12版本中正式支持，它是Go官方提供的包管理解决方法。通过Go Modules，我们可以不必将项目放置到GOPATH上。Go Modules和传统的GOPATH不同，不需要包含三个子目录，一个源代码目录，甚至是空目录都可以作为Module，只要其中包含go.mod文件；
通过Go Modules可以轻易地进行一个包的依赖管理和版本控制，go build和go install将自动使用go.mod中的依赖关系，减少了GOPATH管理时的复杂性。


2、学习笔记
先上资源地址：https://github.com/hoanhan101/ultimate-go
学习笔记大致分为：
五大部分
1.Go语言设计理念，强调了理解代码本身，学写代码其实和学一门自然语言有相似之处，需要阅读好的坏的代码加深语感，而这块是很多程序员缺乏的。
2.语言机制（Language Mechanics），包括Go语言的句法、数据结构、解耦三部分，每个分别有多个细分介绍。
3.软件设计（Software Design），内容细化到分组类型解耦过程、界面转换、界面污染，mock过程和常见雷区。
4.研究并发性（Concurrency），也就是在Go协程（Goroutine）、数据竞赛、多个channel和不用模式语境和模式下的操作过程。
5.测试和分析（Testing and Profiling），写到了基本单元测试、表测试、自测试等发测试方法，以及常见的标准等。此外，还有各种包（Packages），作者表示这部分还在完善中。

Go语言资源大汇总
Go的热度不是盖的，这份资源一出，Hackernews上的讨论区就开始活跃了，不少网友继续贡献资源、分享心得，也有一些对这份学习笔记的评价。
用户@olah_1推荐了一个Go语言课程，表示：
Learn Go with Tests是我经历过的最好的编程语言课程。
地址：https://github.com/quii/learn-go-with-tests
自带中文版资源：
用户@plinkplonk建议，如果搞不明白Go是什么，可以去OReilly上资料：
http://shop.oreilly.com/product/0636920046516.do
用户@ValentineC推荐了一个GitHub上5000+star的课程培训，来自上面提到的机构Ardan Labs：
https://github.com/ardanlabs/gotraining
还有网友贡献自制学习思维导图：
地址：https://github.com/dzyanis/roadmap
也有人推荐学习如何在没有框架的情况下用Go语言编写web app：
https://github.com/thewhitetulip/web-dev-golang-anti-textbook
据说，这份资料是由有15年编程经验人完成，现在免费开放。
有了这么多过来人的建议，还怕学不好么~
传送门
GitHub地址：
https://github.com/hoanhan101/ultimate-go
hackernews讨论区：
https://news.ycombinator.com/item?id=20701671
— 完 —
https://www.zhihu.com/question/23486344


Log包线程安全吗？
Golang的标准库提供了log的机制，但是该模块的功能较为简单（看似简单，其实他有他的设计思路）。在输出的位置做了线程安全的保护。


Golang中除了加Mutex锁以外还有哪些方式安全读写共享变量？
Golang中Goroutine 可以通过 Channel 进行安全读写共享变量。
