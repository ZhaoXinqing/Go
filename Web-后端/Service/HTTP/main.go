
// HTTP协议：
// 超文本传输协议（HTTP，HyperText Transfer Protocol)，
// 是互联网上应用最为广泛的一种网络传输协议，所有的WWW文件都必须遵守这个标准。
// 设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

// HTTP服务端
// net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据。
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

// HTTP客户端
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

