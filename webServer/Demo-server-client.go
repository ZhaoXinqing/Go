package webServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUrl(URL string){
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

