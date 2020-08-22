// // 程序的处理流程：
// // 	1. 监听端口
// // 	2. 接收客户端请求建立链接
// // 	3. 创建goroutine处理链接

// // 处理函数
// func process(conn net.Conn) {
// 	defer conn.Close() // 关闭连接
// 	for {
// 		reader := bufio.NewReader(conn)
// 		var buf [128]byte
// 		n, err := reader.Read(buf[:]) // 读取数据
// 		if err != nil {
// 			fmt.Println("连接客户端失败，错误信息：", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到client端发来的数据：", recvStr)
// 		conn.Write([]byte(recvStr)) //发送数据
// 	}
// }

// func main() {
// 	listen, err := net.Listen("tcp", "127.0.0.0:20000")
// 	if err != nil {
// 		fmt.Println("listen failed：", err)
// 		return
// 	}
// 	for {
// 		conn, err := listen.Accept() // 建立连接
// 		if err != nil {
// 			fmt.Println("accept failed,err:", err)
// 			continue
// 		}
// 		go process(conn) // 启动一个goroutine处理连接
// 	}
// }

// // ## TCP客户端 ##
// // 一个TCP客户端进行TCP通信的流程如下：
// // 1. 建立与服务端的链接
// // 2. 进行数据收发
// // 3. 关闭链接

// //

// // 使用go语言的net包实现TCP客户端代码
// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.0:20000")
// 	if err != nil {
// 		fmt.Println("连接失败，错误：", err)
// 		return
// 	}
// 	defer conn.Close() // 关闭连接
// 	inputReader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := inputReader.ReadString('\n') // 读取用户输入
// 		inputInfo := strings.Trim(input, "\r\n")
// 		if strings.ToUpper(inputInfo) == "Q" { // 如果输入Q就退出
// 			return
// 		}
// 		_, err = conn.Write([]byte(inputInfo)) // 发送数据
// 		if err != nil {
// 			return
// 		}
// 		buf := [512]byte{}
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Println("recv failed, err：", err)
// 			return
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// // Encode 将消息编码
// func Encode(message string)([]byte ,error)  {
//     // 读取消息的长度，转换成int32类型（占4个字节）
//     var length = int32(len(message))
//     var pkg = new(bytes.Buffer)
//     //写入消息头
//     err := binary.Write(pkg,binary.LittleEndian,length)
//     if err != nil{
//         return nil,err
//     }
//     //写入消息实体
//     err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
//     if err != nil{
//         return nil,err
//     }
//     return pkg.Bytes(),nil
// }

// // Decode 消息解码
// func Decode(reader *bufio.Reader)(string,error)  {
//     //读取消息长度
//     lengthByte,_ := reader.Peek(4) //读取前4个字节数据
//     lengthBuff := bytes.NewBuffer(lengthByte)
//     var length int32
//     err := binary.Read(lengthBuff,binary.LittleEndian,&length)
//     if err != nil{
//         return "",err
//     }
//     // Buffered返回缓冲中现有的可读取的字节数。
//     if int32(reader.Buffered()) < length+4{
//         return "",err
//     }
//     //读取真正的消息数据
//     pack := make([]byte,int(4+length))
//     _,err = reader.Read(pack)
//     if err != nil{
//         return "",err
//     }
//     return string(pack[4:]),nil
// }

// // server端：
// func process(conn net.Conn)  {
//     defer conn.Close()
//     reader := bufio.NewReader(conn)
//     for {
//         msg,err := proto.Decode(reader)
//         if err == io.EOF{
//             return
//         }
//         if err != nil{
//             fmt.Println("decode 失败,err",err)
//             return
//         }
//         fmt.Println("收到client数据:",msg)
//     }
// }
// func main()  {
//     listen,err := net.Listen("tcp","127.0.0.1:8888")
//     if err != nil{
//         fmt.Println("监听失败,err",err)
//         return
//     }
//     defer listen.Close()
//     for {
//         conn,err := listen.Accept()
//         if err != nil{
//             fmt.Println("接受失败,err",err)
//             continue
//         }
//         go process(conn)
//     }
// }

// // client端:
// func main()  {
//     conn,err := net.Dial("tcp","127.0.0.1:8888")  // 监听
//     if err != nil{
//         fmt.Println("dial失败，err",err)
//         return
//     }
//     defer conn.Close()		// 关闭连接
//     for i:=0;i<20;i++{
//         msg := "Hello Ares!"
//         data,err := proto.Encode(msg)		// 解码
//         if err != nil{
//             fmt.Println("encode失败，err",err)
//             return
//         }
//         conn.Write(data)		写入
//     }
// }

// // net包实现的UDP服务端：

// // UDP协议（User Datagram Protocol）：用户数据报协议，
// // 是开放式系统互联，参考模型中一种无连接的传输层协议，
// // 不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，
// // 但是UDP协议的实时性比较好，通常用于视频直播相关领域。

// func main() {
// 	listen, err := net.ListenUDP("udp", &net.UDPAddr{
// 		IP:   net.IPv4(0, 0, 0, 0),
// 		Port: 30000,
// 	})
// 	if err != nil {
// 		fmt.Println("listen failed, err:", err)
// 		return
// 	}
// 	defer listen.Close()

// 	for {
// 		var data [1024]byte
// 		n, addr, err := listen.ReadFromUDP(data[:]) // 接受数据
// 		if err != nil {
// 			fmt.Println("read udp failed, err:", err)
// 			continue
// 		}
// 		fmt.Println("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
// 		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
// 		if err != nil {
// 			fmt.Println("Write to udp failed, err:", err)
// 			continue
// 		}
// 	}
// }

// // net包实现的UDP客户端：

// func main() {
// 	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
// 		IP:   net.IPv4(0, 0, 0, 0),
// 		Port: 30000,
// 	})
// 	if err != nil {
// 		fmt.Println("连接服务器失败，错误：", err)
// 		return
// 	}
// 	defer socket.Close()
// 	sendData := []byte("hello server")
// 	_, err = socket.Write(sendData) // 发送数据
// 	if err != nil {
// 		fmt.Println("发送数据失败，err：", err)
// 		return
// 	}
// 	data := make([]byte, 4096)
// 	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
// 	if err != nil {
// 		fmt.Println("接受数据失败，err", err)
// 		return
// 	}
// 	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
// }
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, bengbeng")
}
