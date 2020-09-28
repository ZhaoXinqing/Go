package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Info struct {
	Name string      `json:"name"`
	Path string      `json:"path"`
	Type interface{} `json:"type"`
	Size interface{} `json:"size"`
	Date interface{} `json:"data"`
}

// 上传文件
func Createfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 创建路径
	dirname := params.ByName("dir")
	path := "/rtfdata/third_party/" + dirname

	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("该路径不存在", err)
		// 创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("创建路径失败", err)
		} else {
			fmt.Println("创建路径成功")
		}
	}

	//因为上传文件的类型是multipart/form-data 所以不能使用 r.ParseForm(), 这个只能获得普通post
	r.ParseMultipartForm(32 << 20) //上传最大文件限制32M

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err, "--------1------------") //上传错误
		return
	}
	defer file.Close()

	//创建上传的目的文件
	f, err := os.OpenFile(path+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建文件失败: %s", err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, file) //拷贝文件
}

// 查看目录
func Listfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	infos := make([]Info, 0)
	dir := params.ByName("dir")
	path := "/rtfdata/third_party/" + dir

	files, err := ioutil.ReadDir(path) //读取目录下文件
	if err != nil {
		fmt.Println("文件不存在", err)
		a := "空目录"
		writeResponse(w, a)

	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("dir")
		} else {
			info := Info{file.Name(), path, file.Mode(), file.Size(), file.ModTime()}
			infos = append(infos, info)
		}
	}
	writeResponse(w, infos)
}

// 下载文件
func Getfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dir := params.ByName("dir")
	name := params.ByName("name")
	path := "/rtfdata/third_party/" + dir + "/" + name

	_, err := os.Stat(path)
	if err != nil {
		writeResponse(w, "文件不存在")
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		return
	}
}

// 删除文件
func Deletefile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dir := params.ByName("dir")
	name := params.ByName("name")
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

// 清空目录
func Deletedir(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dir := params.ByName("dir")
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

func loadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	url := p.ByName("url")
	file, err := os.Open("./" + url)
	if err != nil {
		httputil.WriteErrorResponse(w, "文件不存在")
		return
	}
	defer file.Close()
	buff, _ := ioutil.ReadAll(file)
	w.Write(buff)
}
