## socket编程 ##
Socket是BSD UNIX的进程通信机制，通常也称作”套接字”，用于描述IP地址和端口，是一个通信链的句柄。Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序员可以用它们来开发TCP/IP网络上的应用程序。电脑上运行的应用程序通常通过”套接字”向网络发出请求或者应答网络请求。
Socket是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，Socket其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在Socket后面，对用户来说只需要调用Socket规定的相关函数，让Socket去组织符合指定的协议数据然后进行通信。
## GO语言实现传输层通信 ##
一个TCP服务端可以同时连接很多个客户端，Go语言中创建多个goroutine实现并发非常方便和高效，所以可以每建立一次链接就创建一个goroutine去处理。
## TCP服务端 ##
程序的处理流程：

1. 监听端口
1. 接收客户端请求建立链接
1. 创建goroutine处理链接

// TCP server 端

    func process(conn net.Conn) {
        defer conn.Close()  // 关闭连接
        for {
            reader := buffio.NewReader(conn)
            var buf [128]byte
            n,err := reader.Read{buf[:]}  // 读取数据
            if err != nil {
                fmt.Println("连接客户端失败，错误信息："，err)
            }
            recvStr := string{buf[:n]}
            fmt.Println("收到客户端信息：", recvStr)
            conn.Write([]byte{recvStr})  //发送数据
        }
    }
    
    func main() {
        listen, err := net.Listen("tcp","127.0.0.0:8888")
        if err != nil {
            fmt.Println("监听失败，错误：", err)
            return
        }
        for {
            conn, err := listen.Accept() // 建立连接
            if err != nil {
                fmt.Println("建立连接失败，错误：", err)
                continue
            }
            go process(conn)  // 启动一个goroutine处理连接
        }
    }


## TCP客户端 ##
一个TCP客户端进行TCP通信的流程如下：

1. 建立与服务端的链接
1. 进行数据收发
1. 关闭链接


// tcp客户端

    func main() {
        conn, err := net.Dial("tcp","127.0.0.0:8888")
        if err != nil {
            fmt.Println("连接失败，错误：", err)
            return
        }
        defer conn.Close()
        inputReader := bufio.NewReader(os.Stdout)
        for {
            input,_ := inputReader.ReadString('\n')
            inputInfo := string.Trim(input,"\r\n")
            if string.ToUpper(inputInfo) == "q" {
                return // 如果输入q就退出
            }
            _, err = conn.Write([]byte(inputInfo))  // 发送数据
            if err := nil{
                return
            }
            buf := [512]byte{}
            n, err != nil{
                fmt.Println("接受失败，错误：", err)
                return
            }
            fmt.Println(astring(buf[:n]))
        }
    }

## UDP服务端 ##

UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。

//服务端

    func main()  {
        listen,err := net.ListenUDP("udp",&net.UDPAddr{
            IP:net.IPv4(0,0,0,0),
            Port:8888,
        })
        if err != nil{
            fmt.Println("监听失败，错误：",err)
            return
        }
        defer listen.Close()
        for {
            var data [1024]byte
            n,addr,err := listen.ReadFromUDP(data[:])
            if err != nil{
                fmt.Println("接收udp数据失败，错误：",err)
                continue
            }
            fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
            _ ,err = listen.WriteToUDP(data[:n],addr)   //发送数据
            if err != nil{
                fmt.Println("发送数据失败，错误：",err)
                continue
            }
        }
    }

UDP客户端
//客户端

    func main()  {
        socket,err := net.DialUDP("udp",nil,&net.UDPAddr{
            IP:net.IPv4(0,0,0,0),
            Port:8888,
        })
        if err != nil{
            fmt.Println("连接服务器失败，错误：",err)
            return
        }
        defer socket.Close()
        sendData := []byte("hello world!")
        _,err = socket.Write(sendData)
        if err != nil{
            fmt.Println("发送数据失败，错误：",err)
            return
        }
        data := make([]byte,4096)
        n,remoteAddr,err := socket.ReadFromUDP(data)
        if err != nil{
            fmt.Println("接受数据失败，错误：",err)
            return
        }
        fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
    }

## HTTP客户端和服务端 ##

**HTTP协议**
超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

**HTTP服务端**

	net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据。
	// http server
	func sayHi(w http.ResponseWriter,r *http.Request)  {
	    fmt.Fprintln(w,"你好，ares！")
	}
	func main()  {
	    http.HandleFunc("/",sayHi)
	    err := http.ListenAndServe(":8888",nil)
	    if err != nil{
	        fmt.Println("Http 服务建立失败，err：",err)
	        return
	    }
	}

HTTP客户端

	func main() {
	    resp, err := http.Get("https://www.baidu.com/")
	    if err != nil {
	        fmt.Println("get failed, err:", err)
	        return
	    }
	    defer resp.Body.Close()
	    body, err := ioutil.ReadAll(resp.Body)
	    fmt.Printf("%T\n",body)
	    fmt.Println(string(body))
	}

## TCP粘包 ##

在socket网络程序中，TCP和UDP分别是面向连接和非面向连接的。因此TCP的socket编程，收发两端（客户端和服务器端）都要有成对的socket，因此，发送端为了将多个发往接收端的包，更有效的发到对方，使用了优化方法（Nagle算法），将多次间隔较小、数据量小的数据，合并成一个大的数据块，然后进行封包。这样，接收端，就难于分辨出来了，必须提供科学的拆包机制。

对于UDP，不会使用块的合并优化算法，这样，实际上目前认为，是由于UDP支持的是一对多的模式，所以接收端的skbuff(套接字缓冲区）采用了链式结构来记录每一个到达的UDP包，在每个UDP包中就有了消息头（消息来源地址，端口等信息），这样，对于接收端来说，就容易进行区分处理了。所以UDP不会出现粘包问题。

**粘包产生原因**

1.发送端需要等缓冲区满才发送出去，造成粘包

2.接收方不及时接收缓冲区的包，造成多个包接收

具体点：

1. 发送方引起的粘包是由TCP协议本身造成的，TCP为提高传输效率，发送方往往要收集到足够多的数据后才发送一包数据。若连续几次发送的数据都很少，通常TCP会根据优化算法把这些数据合成一包后一次发送出去，这样接收方就收到了粘包数据。
1. 接收方引起的粘包是由于接收方用户进程不及时接收数据，从而导致粘包现象。这是因为接收方先把收到的数据放在系统接收缓冲区，用户进程从该缓冲区取数据，若下一包数据到达时前一包数据尚未被用户进程取走，则下一包数据放到系统接收缓冲区时就接到前一包数据之后，而用户进程根据预先设定的缓冲区大小从系统接收缓冲区取数据，这样就一次取到了多包数据。

**解决办法**

出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。	自定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

	// Encode 将消息编码
	func Encode(message string)([]byte ,error)  {
	    // 读取消息的长度，转换成int32类型（占4个字节）
	    var length = int32(len(message))
	    var pkg = new(bytes.Buffer)
	    //写入消息头
	    err := binary.Write(pkg,binary.LittleEndian,length)
	    if err != nil{
	        return nil,err
	    }
	    //写入消息实体
	    err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	    if err != nil{
	        return nil,err
	    }
	    return pkg.Bytes(),nil
	}
	
	// Decode 消息解码
	func Decode(reader *bufio.Reader)(string,error)  {
	    //读取消息长度
	    lengthByte,_ := reader.Peek(4) //读取前4个字节数据
	    lengthBuff := bytes.NewBuffer(lengthByte)
	    var length int32
	    err := binary.Read(lengthBuff,binary.LittleEndian,&length)
	    if err != nil{
	        return "",err
	    }
	    // Buffered返回缓冲中现有的可读取的字节数。
	    if int32(reader.Buffered()) < length+4{
	        return "",err
	    }
	    //读取真正的消息数据
	    pack := make([]byte,int(4+length))
	    _,err = reader.Read(pack)
	    if err != nil{
	        return "",err
	    }
	    return string(pack[4:]),nil
	}
	
	// server端：
	func process(conn net.Conn)  {
	    defer conn.Close()
	    reader := bufio.NewReader(conn)
	    for {
	        msg,err := proto.Decode(reader)
	        if err == io.EOF{
	            return
	        }
	        if err != nil{
	            fmt.Println("decode 失败,err",err)
	            return
	        }
	        fmt.Println("收到client数据:",msg)
	    }
	}
	func main()  {
	    listen,err := net.Listen("tcp","127.0.0.1:8888")
	    if err != nil{
	        fmt.Println("监听失败,err",err)
	        return
	    }
	    defer listen.Close()
	    for {
	        conn,err := listen.Accept()
	        if err != nil{
	            fmt.Println("接受失败,err",err)
	            continue
	        }
	        go process(conn)
	    }
	}
	
	// client端:
	func main()  {
	    conn,err := net.Dial("tcp","127.0.0.1:8888")
	    if err != nil{
	        fmt.Println("dial失败，err",err)
	        return
	    }
	    defer conn.Close()
	    for i:=0;i<20;i++{
	        msg := "Hello Ares!"
	        data,err := proto.Encode(msg)
	        if err != nil{
	            fmt.Println("encode失败，err",err)
	            return
	        }
	        conn.Write(data)
	    }
	}
