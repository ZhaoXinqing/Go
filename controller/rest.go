package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MiniProgramUserInfo 用户信息
type MiniProgramUserInfo struct {
	ID       int64
	UserName string
	Phone    string
	Avatar   string
	Name     string
}

// FromHeaderToUserInfo 添加header的token 返回用户信息
func FromHeaderToUserInfo(token string) (dataInfo *MiniProgramUserInfo, err error) {
	client := &http.Client{}
	url := "http://192.168.2.135:8092/penelope/service/v1/auth"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", token)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(request)
	defer response.Body.Close()
	bodyContent, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(bodyContent, &dataInfo)
	if err != nil {
		fmt.Print(err)
	}
	return
}
