package main

import (
	"fmt"
	"net/http"
)

// net/http包,专门用来处理 HTTP 协议的数据

// server
func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "你好，iris!")
}

func main() {
	http.HandleFun("/", sayHi)
	err := http.ListenAndServer(":8080", nil)
	if err != nil {
		fmt.Println("Http 服务建立失败，err: ", err)
		return
	}
}

// client
/*
func main() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("%T\n", body)
	fmt.Println(string(body))
}
*/
