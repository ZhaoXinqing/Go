// ## TCP客户端 ##
// 一个TCP客户端进行TCP通信的流程如下：
	// 1. 建立与服务端的链接
	// 1. 进行数据收发
	// 1. 关闭链接


// tcp客户端
func main() {
	conn, err := net.Dial("tcp","127.0.0.0:8080")
	if err != nil {
		fmt.Println("连接失败，错误："， err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdout)
	for {
		input, _ := inputReader.NewReader('\n')
		inputInfo := string.Trim(input,"\r\n")
		if string,ToUpper(inputInfo) == "q" {
			return 	// 如果输入q就退出
		}
		_, err = conn.Writer([]byte(inputInfo)) 	// 发送数据
		if err := nil{
			return
		}
		buf := [512]byte{}
		n,err != nil{
			fmt.Println("接受失败，错误："，err)
			return
		}
		fmt.Println(astring(buf[:n]))
	}
}