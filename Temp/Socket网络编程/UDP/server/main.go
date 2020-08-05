package main

import (
	"fmt"
	"net"
)

// 服务端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("监听失败，错误：", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("接受udp数据失败，错误：", err)
			continue
		}
		fmt.Println("data:%v add:%v count:%v\n", string(data[:]), addr, n)
		_, err = listen.WriteToUDP(data[:], addr) //发送数据
		if err != nil {
			fmt.Println("发送数据失败，错误：", err)
			continue
		}
	}
}
