IO简介:
    - Unix（like）中，一切皆文件。Socket、FIFO、管道、终端都是文件，一切都是流。在信息交换的过程中，实际都是对这些流进行的数据收发操作，简称I/O操作（系统调用read、write）。而流有很多，于是就用文件描述符（fd）来区分具体是哪个流。例如，我们创建了一个网络io(socket)，系统调用会返回一个fd，对socket的任何操作都是对这个fd的操作（隐隐包含着一种分层与抽象的思想）。
    
    - 同步与异步：
        - 是一种通信机制，涉及到调用方和被调用方（针对应用程序与内核而言）。
            - 同步过程中，进程触发IO操作并等待（阻塞）或者轮询的（非阻塞）去查看IO操作是否完成；
            - 异步过程中，进程触发IO操作以后，直接返回，做自己的事情，IO交给内核来处理，完成后内核通知进程IO完成。
        - 同步和异步关注的是程序之间的协作关系。同步分为阻塞和非阻塞，异步则只有非阻塞。

    - 阻塞和非阻塞：
        - 是一种调用机制，只涉及到调用方（针对单个进程的执行状态），关注的是IO操作的执行。
        - 调用方等待IO操作完成后返回则为阻塞；调用方无需等待IO操作完成便返回则为非阻塞，在非阻塞的情况下，调用方常常需要主动去check，获得IO的操作结果。    
    
    - 五种IO模型：
        - IO发生时（以网络IO read为例)涉及到两个系统对象和两个阶段，IO模型的区别就是在这两个阶段上的差异。
            - 两个系统对象：一个是调用这个IO的process（进程）或者thread（线程），另一个就是系统内核；
            - 两个阶段：1、调用方等待数据准备，2、内核将数据从内核copy到process中。 

        1、blocking（阻塞）IO：
            - 当调用一个系统调用read时，kernel就开始IO的第一个阶段：准备数据，对于network IO来说，很多时候数据一开始还没有到达（一个TCP包没有接收完整），这个时候kernel就要等待足够的数据到来（这也和缓存IO还是非缓存IO有关，一般都是缓存IO）。而在用户进程这边，整个进程会被阻塞。当kernel一直等到数据准备好了，就会将数据从系统内存copy到用户内存，然后kernel返回结果，用户进程才接触block状态，重新运行起来。blocking IO的特点就是两个阶段都被block。
        2）non-blocking（非阻塞）IO：
            - 用户进程发出read操作时，如果kernel中的数据还没有准备好，那么它并不会block用户进程，而是立刻返回一个error。用户进程判断结果是error时，他就知道数据还没有准备好，于是就再次发送read调用。直到kernel中数据准备好了后，并且再次接收到system call后，将数据copy到用户内存，但是这种模型效率很低。
        3）multiplexing（多路复用）IO：
            - 也称事件驱动（event driven IO），使用select/epoll通过一个thread可以同时处理多个socket的IO，其基本原理就是select/epoll会不断轮询所负责的socket，当某个socket有数据到达了，就通知用户进程。当用户进程调用了select/epoll，整个进程就会被block，同时kernel会observe所有select/epoll负责的socket，任何一个socket中数据准备好之后，select/epoll就会返回，这时用户进程再调用read system call，将数据从kernel copy进用户内存。
            - multiplexing io因为有两个system call，甚至比blocking io更差一些，优势是它可以处理多个connection。所以使用select/epoll的web server不一定处理速度很快，他只是能处理更多连接。
        4）asynchronous(异步)IO
            - 用户进程发起read操作之后，就可以去做其它的事。而kernel收到asynchronous read之后，会立刻返回，所以对用户进程不会产生任何block。当数据准备完成，kernel会等将数据拷贝到用户内存，并给用户进程发送一个signal，告诉它read操作完成了。
        5）信号驱动io模型：
            - 一般用于UDP中，当用户线程发起一个IO请求操作，会给对应的socket注册一个信号函数，然后用户线程会继续执行，当内核数据就绪时会发送一个信号给用户线程，用户线程接收到信号之后，便在信号函数中调用IO读写操作来进行实际的IO请求操作。这个，对TCP套接口几乎是没用的，原因是该信号产生得过于频繁，并且该信号的出现并没有告诉我们发生了什么事情。
        五种io模型详解：https://blog.csdn.net/ocean_fan/article/details/79622956
      
io多路复用：
    - 如何使用:
        - 利用select、poll、epoll可以同时监察多个流的IO事件的能力，在空闲的时候会把当前线程阻塞，当有一个或多个流由IO事件发生时，就从阻塞态中唤醒，于是程序就会轮询一遍所有的流（epoll只是轮询那些真正发出了事件的流），并且依次顺序的处理就绪的流。
            - Select：
                -服务端一直轮询、监听如果有客户端链接上来就创建一个连接放到数组A中，继续轮询这个数组，如果轮询的过程中有客户端发生IO事件就去处理；select'只能监视1024个连接（一个进程只能创建1024个文件）；而且存在线程安全的问题。
            - poll：
                - 在select做了许多修复，比如不限制检测的连接数；但是也有线程安全问题；
            - epoll：
                - 也是检测IO事件，但是如果发生io事件，它会告诉你是哪个连接发生了事件，就不用再轮询访问。而且它是线程安全的，但是只有linux平台支持；
            - 对比：
                - select, poll是为了解決同时大量IO的情況（尤其网络服务器），但是随着连接数越多，性能越差。epoll是select和poll的改进方案，在 linux 上可以取代 select 和 poll，可以处理大量连接的性能问题；
    
    - Epoll：     
        - 是是Linux下多路复用IO接口select/poll的增强版本，是一种io多路复用技术，可以非常高效的处理数以百万计的Socket句柄，比Select和Poll效率提高了很多。
        
        - 优势：
            - 它能显著提高程序在大量并发连接中只有少量活跃的情况下的系统CPU利用率。
            - epoll除了提供select/poll那种IO事件的水平触发（Level Triggered）外，还提供了边缘触发（Edge Triggered），这就使得用户空间程序有可能缓存IO状态，减少epoll_wait/epoll_pwait的调用，提高应用程序效率。
        
        - epoll提供的三个函数：
            - epoll_create(int size):
                - 建立一个epoll对象，并传回它的id。
            - epoll ctl(int epfd, int op, int fd, struct epoll, epoll_event *event):
                - 时间注册函数，将需要监听的事件和需要监听的fd交给epoll对象。
            - epoll_wait(int epfd, struct epoll_event *event, int maxevents, int timeout):
                - 等待注册的事件被处罚或者timeout发生。
        
        - Epoll的两种触发模式：
            - 水平模式LT：
                - LT就是与selet和poll类似，当被监控的文件描述符上有可读写事件发生时，epoll_wait()会通知处理程序去读写。如果这次没有把数据一次性全部读写完（如读写缓冲区太小），那么下次调用epoll_wait() 时，它还会通知你在上次没读写完的文件描述符上继续读写；
            - 边缘模式ET：
        
        - 原理：
            - 获取事件的时候，它无须遍历整个被侦听的描述符集，只要遍历那些被内核IO事件异步唤醒而加入Ready队列的描述符集合就行了。


    【一篇入魂】网络编程中的五种IO模型](https://blog.csdn.net/jiaodaguan/article/details/104000052)<br>

    - 10种网络io模型及使用：
        - 单进程同步：NTP
        - 多线程同步：Natty
        - 单线程异步：Redis
        - 半同步半异步：Natty
        - 多进程同步：fastcgi
        - 多线程异步：memcached
        - 多进程异步：nginx
        - 每请求每进程（线程）：Apache/CGI
        - 微进程框架：erlang/go/lua
        - 协程框架：libco/ntyco/libgo

    - 深入理解阻塞
        - 因为一个线程只能处理一个socket的IO，如果想同时处理多个，可以用非阻塞忙轮询的方法，伪代码是这样的：
            for{
                for _,v := range []streams{
                    if v has data
                    read until unavailable
                }
            }
        - 把流（stream）从头到尾读一遍就能处理了，可是这样效率很低，要是所有流都没有IO事件，就浪费了CPU的资源。为了避免CPU空转，不让这个线程去检查流是否有IO事件，而是引进一个代理（起初是select，后来是epoll），它可以同时observe许多stream事件，如果没有事件，代理就阻塞，线程就不会去挨个check了。伪代码：
            for{
                select([]streams) 
                for _,v := range []streams{
                    if v has data
                    read until unavailable
                }
            }
        - 可即便这样，依旧要去做循环，因为select只是告诉线程有IO事件发生，可并没有告诉线程是哪个fd（或者多个），所以select的复杂度是O(n)。epoll（event poll）就解决了这个问题，所以epoll是事件驱动的，因为每个事件上关联了fd，复杂度也降到了O(1)。
            for{
                happened_IO := append(happened_IO,epoll_wait(epollfd))
                for _,v := range happened_IO{
                    read or write
                }
            }

