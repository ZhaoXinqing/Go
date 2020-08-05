func main() {
	socket,err := net.DialUDP("udp",nil,&UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:8080,
	})
	if err != nil {
		fmt.Println("连接服务器失败，错误："，err)
		return
	}
	defer socket.Close()
	sendData := []byte("hello world!")
	_,err != socket.Write(sendData)
	if err != nil{
		fmt.Println("发送数据失败，错误：",err)
		return
	}
	data := make([]byte,4096)
	n,remoteAddr,err := socket.ReadFromUDP(data)
	if err != nil{
		fmt.Println("接受数据失败，错误",err)
		return
	}
	fmt.Println("recv:%v addr:%v count:%v\n",string(data[:n]),remoteAddr,n)
}