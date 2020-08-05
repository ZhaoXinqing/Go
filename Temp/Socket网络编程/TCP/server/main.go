// TCP服务端 ##
// 程序的处理流程：
// 	1. 监听端口
// 	1. 接收客户端请求建立链接
// 	1. 创建goroutine处理链接

// TCP server 端
func process(conn net.Conn) {
	defer conn.Close() 	// 关闭连接
	for {
		reader := buffio.NewReader(conn)
		var buf [128]byte
		n, err := raader.Read{buf[:]}	// 读取数据
		if err != nil {
			fmt.Println("连接客户端失败，错误信息："，err)
		}
		recvStr := string{buf[:n]}
		fmt.Println("收到客户端信息："，recvStr)
		conn.Write([]byte{recvStr})	//发送数据
	}
}

func main() {
	listen,err := net.Listen("tcp","127.0.0.0:8080")
	if err != nil {
		fmt.Println("监听失败，错误："err)
		return
	}
	for {
		conn, err := listen.Accept()	// 建立连接
		if err != nil {
			fmt.Println("建立连接失败，错误："，err)
			continue
		}
		go process(conn) 	// 启动一个goroutine处理连接
	}
}
