package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 使用 net 包实现TCP客户端代码
func main() {
	conn, err := net.Dial("tcp", "127.0.0.0:20000")
	if err != nil {
		fmt.Println("连接失败，错误：", err)
		return
	}
	defer conn.Close() //关闭连接

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == Q {
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
