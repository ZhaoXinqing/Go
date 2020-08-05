    1、CSP并发模型：
        - Go语言中实现了两种并发模型，一种是我们熟悉的线程和锁并发模型，它主要依赖于共享内存实现。另一种是Go语言倡导使用的CSP（communicating sequential processes）通信顺序进程模型，是通过Goroutine和Channel来实现的。
        - 通信机制channel，在通信过程中，传数据channel <- data和取数据<-channel必然会成对出现，两个goroutine之间才会实现通信，否则必阻塞，直到另外的goroutine传或者取为止。

    2、MPG线程模型：
        - Go语言的MPG模型属于一种特殊的两级线程模型（用户级线程模型、内核级线程模型），它将CPU、内核线程、线程的关系描述为M、P、G三者的关系；
            - M(Machine)：
                - 一个machine对应一个内核线程，相当于内核线程在Go语言进程中的映射；
            - P(Processor)：
                - 是让我们从N:1调度到M:N调度的重要部分，它的主要用途就是维护了一个goroutine队列（即runqueue），用来执行goroutine，一个processor表示要执行Go代码片段的所必须的上下文环境，可以理解为用户代码逻辑的处理器；
            - G(Goroutine)：
                - 是一种轻量级的用户线程，是对Go语言中代码片段的封装，它包含了栈，指令指针，以及其它对调度goroutine很重要的信息，例如其阻塞的channel。
        - P的最大数量决定了程序的并发规模，且P的最大数量是由程序决定的。可以通过修改环境变量GOMAXPROCS和调用函数runtime.GOMAXPROCS来设定最大值。
