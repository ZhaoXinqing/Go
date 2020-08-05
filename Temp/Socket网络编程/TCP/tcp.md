
## TCP粘包 ##

在socket网络程序中，TCP和UDP分别是面向连接和非面向连接的。因此TCP的socket编程，收发两端（客户端和服务器端）都要有成对的socket，因此，发送端为了将多个发往接收端的包，更有效的发到对方，使用了优化方法（Nagle算法），将多次间隔较小、数据量小的数据，合并成一个大的数据块，然后进行封包。这样，接收端，就难于分辨出来了，必须提供科学的拆包机制。

对于UDP，不会使用块的合并优化算法，这样，实际上目前认为，是由于UDP支持的是一对多的模式，所以接收端的skbuff(套接字缓冲区）采用了链式结构来记录每一个到达的UDP包，在每个UDP包中就有了消息头（消息来源地址，端口等信息），这样，对于接收端来说，就容易进行区分处理了。所以UDP不会出现粘包问题。

粘包产生原因
    1.发送端需要等缓冲区满才发送出去，造成粘包
    2.接收方不及时接收缓冲区的包，造成多个包接收
    具体点：
        - 发送方引起的粘包是由TCP协议本身造成的，TCP为提高传输效率，发送方往往要收集到足够多的数据后才发送一包数据。若连续几次发送的数据都很少，通常TCP会根据优化算法把这些数据合成一包后一次发送出去，这样接收方就收到了粘包数据。
        - 接收方引起的粘包是由于接收方用户进程不及时接收数据，从而导致粘包现象。这是因为接收方先把收到的数据放在系统接收缓冲区，用户进程从该缓冲区取数据，若下一包数据到达时前一包数据尚未被用户进程取走，则下一包数据放到系统接收缓冲区时就接到前一包数据之后，而用户进程根据预先设定的缓冲区大小从系统接收缓冲区取数据，这样就一次取到了多包数据。

解决办法：
    - 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。	自定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

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