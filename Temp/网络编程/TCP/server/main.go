package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP服务端

// 一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
// 因为Go语言中创建多个goroutine实现并发非常方便和高效，
// 所以我们可以每建立一次链接就创建一个goroutine去处理。

// 程序的处理流程：
// 	1. 监听端口
// 	2. 接收客户端请求建立链接
// 	3. 创建goroutine处理链接

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("连接客户端失败，错误信息：", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed：", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
