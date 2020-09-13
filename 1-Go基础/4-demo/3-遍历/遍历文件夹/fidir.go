package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

// 并发遍历
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
