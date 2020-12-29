package webServer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// GetImageFromURL ...
func GetImageFromURL(url, name string) (avatar string, err error) {
	imageName := name + ".jpg"
	savePath := fmt.Sprintf("uploads/avatars/%s", imageName)

	out, err := os.Create(savePath)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)

	_, err = io.Copy(out, bytes.NewReader(pix))
	return savePath, nil
}

// Deletefile 删除文件
func Deletefile(dir, name string) {
	path := "/rtfdata/third_party/" + dir + "/" + name
	//判断文件是否存在，并删除
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("文件不存在", err)
	} else {
		err = os.Remove(path)
		if err == nil {
			return
		}
	}
}

// Deletedir 清空目录
func Deletedir(dir string) {
	path := "/rtfdata/third_party/" + dir
	// 判断目录是否存在，并删除
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("目录不存在", err)
	} else {
		err = os.RemoveAll(path)
		if err != nil {
			return
		}
	}
}
