package Basic

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// 遍历文件夹
func findDir(dir string, num int) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, fi := range fileinfo {
		print(strings.Repeat("\t", num))
		if fi.IsDir() {
			println(`目录：`, fi.Name())
			findDir(dir+`/`+fi.Name(), num+1)
		} else {
			println(`文件：`, fi.Name())
		}
	}
}

func main() {
	dir := `F:\1-财经知识`
	findDir(dir, 0) // 0为要遍历的层级，默认是0
	fmt.Println(time.Now())
}
