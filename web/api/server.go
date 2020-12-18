package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// 将“Hello World"字符串写入到ResponseWriter, 并格式化字符串，将当前请求的路径添加进去
	fmt.Println(w, "Hello world!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	// 启动服务器并监听8000/tcp端口
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// http包如何实现高并发？
// http包中，server每接收一个用户请求，
// 都会生成一个conn链接，并生成一个goroutines来处理对应的conn。
// 所以每个请求都是独立的，相互不阻塞的。

// https://www.jianshu.com/p/c90ebdd5d1a1


// 请求和响应的首部都使用Header类型表示，
// header类型是一个映射(map)类型，表示HTTP首部中多个键值对。

// 请求和响应的主体都由Request结构中的Body表示，
// 这个字段是Reader和Closer接口的结合

// 表单是客户端和服务端进行数据交互的载体。
// 通常，表单由POST请求在请求体中传递的；

// multipart/form-data格式常用于文件上传功能
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 
		r.ParseMultipartForm(1024)
		// 使用r.FormFile获取文件句柄，然后对文件进行存储处理
		file,handler,err := r.FormFile("uploadfile")
		defer file.Close()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, "%v",handler.header)

		// 假设已经有upload目录，存储文件。
		f,err := os.Openfile("../static/upload/"+handler.Filename,os.O_WRONLY|OS.o.CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close
		io.Copy(f,file)
	} else {
		t, _:= template.ParseFiles("upload.html")
		t.Execute(w,nil)
	}
}

func main() {
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		fmt.Println(err)
	}
}