## 一：IO简介 ##
Unix（like）中，一切皆文件。Socket、FIFO、管道、终端都是文件，一切都是流。在信息交换的过程中，实际都是对这些流进行的数据收发操作，简称I/O操作（系统调用read、write）。而流有很多，于是就用文件描述符（fd）来区分具体是哪个流。For example，我们创建了一个socket，系统调用会返回一个fd，对socket的任何操作都是对这个fd的操作（隐隐包含着一种分层与抽象的思想）。

## 二：同步异步、阻塞非阻塞 ##
**同步与异步是一种通信机制，涉及到调用方和被调用方（针对应用程序与内核而言）**。同步过程中，进程触发IO操作并等待（阻塞）或者轮询的（非阻塞）去查看IO操作是否完成；异步过程中，进程触发IO操作以后，直接返回，做自己的事情，IO交给内核来处理，完成后内核通知进程IO完成。同步和异步关注的是程序之间的协作关系。同步分为阻塞和非阻塞，异步则只有非阻塞。

**阻塞和非阻塞是一种调用机制，只涉及到调用方（针对单个进程的执行状态）**，关注的是IO操作的执行。调用方等待IO操作完成后返回则为阻塞；调用方无需等待IO操作完成便返回则为非阻塞，在非阻塞的情况下，调用方常常需要主动去check,获得IO的操作结果。

## 三：深入下阻塞 ##
因为一个线程只能处理一个socket的IO，如果想同时处理多个，可以用非阻塞忙轮询的方法，伪代码是这样的：

    for{
    	for _,v := range []streams{
    		if v has data
    		read until unavailable
    	}
    }
把流（stream）从头到尾读一遍就能处理了，可是这样效率很低，要是所有流都没有IO事件，就浪费了CPU的资源。为了避免CPU空转，不让这个线程去检查流是否有IO事件，而是引进一个代理（起初是select，后来是epoll），它可以同时observe许多stream事件，如果没有事件，代理就阻塞，线程就不会去挨个check了。伪代码：

    for{
    	select([]streams) 
    	for _,v := range []streams{
    		if v has data
    		read until unavailable
    	}
    }
可即便这样，依旧要去做循环，因为select只是告诉线程有IO事件发生，可并没有告诉线程是哪个fd（或者多个），所以select的复杂度是O(n)。epoll（event poll）就解决了这个问题，所以epoll是事件驱动的，因为每个事件上关联了fd，复杂度也降到了O(1)。

    for{
    	happened_IO := append(happened_IO,epoll_wait(epollfd))
    	for _,v := range happened_IO{
      		read or write
    	}
    }

## 四：几种IO模型 ##
IO发生时（以network IO read为例）涉及到两个系统对象:一个是调用这个IO的process（进程）或者thread（线程），以及两个阶段：1、等待数据准备，2、将数据从内核copy到process中。 IO模型的区别就是在这两个阶段上的差异。

### 1）blocking（阻塞）IO ###
当调用一个系统调用read时，kernel就开始IO的第一个阶段：准备数据，对于network IO来说，很多时候数据一开始还没有到达（一个TCP包没有接收完整），这个时候kernel就要等待足够的数据到来（这也和缓存IO还是非缓存IO有关，一般都是缓存IO）。而在用户进程这边，整个进程会被阻塞。当kernel一直等到数据准备好了，就会将数据从系统内存copy到用户内存，然后kernel返回结果，用户进程才接触block状态，重新运行起来。blocking IO的特点就是两个阶段都被block。

### 2）non-blocking（非阻塞）IO ###
当用户进程发出read操作时，如果kernel中的数据还没有准备好，那么它并不会block用户进程，而是立刻返回一个error。从用户进程角度来讲，他发起一个read操作后，并不需要等待，而是马上就得到了一个结果。用户进程判断结果是error时，他就知道数据还没有准备好，于是就再次发送read调用。知道kernel中数据准备好了后，并且再次接收到system call后，将数据copy到用户内存。但是这种模型效率很低，上文“深入下阻塞”有提到。

### 3）multiplexing（多路复用）IO ###
其实就是select/epoll，也称为event driven IO。select/epoll的好处就是一个thread可以同时处理多个socket的IO，其基本原理就是select/epoll会不断轮询所负责的socket，当某个socket有数据到达了，就通知用户进程。当用户进程调用了select/epoll，整个进程就会被block，同时kernel会observe所有select/epoll负责的socket，任何一个socket中数据准备好之后，select/epoll就会返回，这时用户进程再调用read system call，将数据从kernel copy进用户内存。
multiplexing io和blocking io差别不大，还更差一些，因为他有两个system call，但是他的优势是他可以处理多个connection。所以使用select/epoll的web server不一定处理速度很快，他只是能处理更多连接。

### 4）asynchronous（异步）IO ###
用户进程发起read操作之后，立刻就可以开始去做其它的事。而另一方面，从kernel的角度，当它受到一个asynchronous read之后，首先它会立刻返回，所以不会对用户进程产生任何block。然后，kernel会等待数据准备完成，然后将数据拷贝到用户内存，当这一切都完成之后，kernel会给用户进程发送一个signal，告诉它read操作完成了。

主要参考了Richard Stevens的“UNIX® Network Programming Volume 1, Third Edition: The Sockets Networking ”的I/O Models