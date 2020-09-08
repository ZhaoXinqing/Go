package api

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// 获取rtfrep.tgz
func GetRtfrep(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// file write
	fw, err := os.Create("rtfrep.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// 打开文件夹
	dir, err := os.Open("/rtfdata")
	if err != nil {
		panic(nil)
	}
	defer dir.Close()

	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	// 遍历文件列表
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		_, err = io.Copy(w, fr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("tar.gz ok")
}

// 更新 rtfrep.tgz
// 接收 r.body，
func UpdataRtfrep(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// 解析传输包，获取上传内容 file
	r.ParseMultipartForm(32 << 20) //上传最大文件限制32M

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err, "--------2------------") //上传错误
		return
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// 显示文件
		fmt.Println(h.Name)
		// 打开文件
		fw, err := os.OpenFile("/rtfdata", os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("un tar.gz ok")
}
