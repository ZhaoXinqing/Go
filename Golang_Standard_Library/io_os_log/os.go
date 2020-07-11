package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("dir:", dir)
	err := os.Chdir("d:/project/test2")
	dir, _ = os.Getwd()
	fmt.Println("dir:", dir)

	//参数不区分大小写
	//不存在环境变量就返回空字符串 len(path) = 0
	path := os.Getenv("gopath")
	fmt.Println(path)

	//返回有效group id
	egid := os.Getegid()
	fmt.Println("egid:", egid)

	//返回有效UID
	euid := os.Geteuid()
	fmt.Println("euid:", euid)

	gid := os.Getgid()
	fmt.Println("gid:", gid)

	uid := os.Getuid()
	fmt.Println("uid:", uid)

	//err:getgroups: not supported by windows
	g, err := os.Getgroups()
	fmt.Println(g, "error", err)

	pagesize := os.Getpagesize()
	fmt.Println("pagesize:", pagesize)

	ppid := os.Getppid()
	fmt.Println("ppid", ppid)

	//filemode, err := os.Stat("main.go")
	//不存在文件返回GetFileAttributesEx test2: The system cannot find the file specified.
	filemode, err := os.Stat("main.go")
	if err == nil {
		fmt.Println("Filename:", filemode.Name())
		fmt.Println("Filesize:", filemode.Size())
		fmt.Println("Filemode:", filemode.Mode())
		fmt.Println("Modtime:", filemode.ModTime())
		fmt.Println("IS_DIR", filemode.IsDir())
		fmt.Println("SYS", filemode.Sys())
	} else {
		fmt.Println("os.Stat error", err)
	}

	//Chmod is not supported under windows.
	//在windows变化是这样子的 -rw-rw-rw- => -r--r--r--
	err = os.Chmod("main.go", 7777)
	fmt.Println("chmod:", err)
	filemode, err = os.Stat("main.go")
	fmt.Println("Filemode:", filemode.Mode())

	//access time modification time
	err = os.Chtimes("main.go", time.Now(), time.Now())
	fmt.Println("Chtime error:", err)

	//获取全部的环境变量
	data := os.Environ()
	for _, val := range data {
		fmt.Println(val)
	}
	fmt.Println("---------end---environ----------------------")

	mapping := func(s string) string {
		m := map[string]string{"xx": "sssssssssssss",
			"yy": "ttttttttttttttt"}
		return m[s]
	}
	datas := "hello $xx blog address $yy"
	//这个函数感觉还蛮有用处
	expandStr := os.Expand(datas, mapping)
	fmt.Println(expandStr)
	datas = "GOBIN PATH $gopaTh" //不区分大小写
	fmt.Println(os.ExpandEnv(datas))

	hostname, err := os.Hostname()
	fmt.Println("hostname:", hostname)

	_, err = os.Open("WWWW.XX")
	if err != nil {
		fmt.Println(os.IsNotExist(err))
		fmt.Println(err)
	}

	f, err := os.Open("WWWW.XX")
	if err != nil && !os.IsExist(err) {
		fmt.Println(f, "not exist")
	}

	//windows 下两个都是true
	fmt.Println(os.IsPathSeparator('/'))
	fmt.Println(os.IsPathSeparator('\\'))
	fmt.Println(os.IsPathSeparator('.'))

	//判断返回的error 是否是因为权限的问题
	//func IsPermission(err error) bool

	// not supported by windows
	err = os.Link("main.go", "newmain.go")
	if err != nil {
		fmt.Println(err)
	}

	var pathSep string
	if os.IsPathSeparator('\\') {
		pathSep = "\\"
	} else {
		pathSep = "/"
	}
	fmt.Println("PathSeparator:", pathSep)
	//MkdirAll 创建的是所有下级目录，如果没有就创建他
	//Mkdir 创建目录，如果是多级目录遇到还未创建的就会报错
	err = os.Mkdir(dir+pathSep+"md"+pathSep+"md"+pathSep+"md"+pathSep+"md"+pathSep+"md", os.ModePerm)
	if err != nil {
		fmt.Println(os.IsExist(err), err)
	}

	err = os.RemoveAll(dir + "md\\md\\md\\md\\md")
	fmt.Println("removall", err)

	//rename 实际上通过movefile来实现的
	err = os.Rename("main.go", "main1.go")

	f1, _ := os.Stat("main.go")
	f2, _ := os.Stat("main1.go")
	if os.SameFile(f1, f2) {
		fmt.Println("the sanme")
	} else {
		fmt.Println("not same")
	}

	//os.Setenv 这个函数是设置环境变量的很简单
	evn := os.Getenv("WD_PATH")
	fmt.Println("WD_PATH:", evn)
	err = os.Setenv("WD_PATH", "D:/project")
	if err != nil {
		fmt.Println(err)
	}

	tmp, _ := ioutil.TempDir(dir, "tmp")
	fmt.Println(tmp)
	tmp = os.TempDir()
	fmt.Println(tmp)

	cf, err := os.Create("golang.go")
	defer cf.Close()
	fmt.Println(err)
	fmt.Println(reflect.ValueOf(f).Type())

	of, err := os.OpenFile("golang.goss", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer of.Close()
	fmt.Println("os.OpenFile:", err)

	oof, err := os.Open("golang.goss")
	defer oof.Close()
	fmt.Println("os.Open file:", oof.Fd())
	fmt.Println("os.Open err:", err)
	oof.Close()

	r, w, err := os.Pipe()
	w.Write([]byte("1111"))
	var buf = make([]byte, 4)
	r.Read(buf)
	fmt.Println(buf)
	w.Write([]byte("2222"))
	r.Read(buf) // 如果没有调用w.Write(),r.Read()就会阻塞
	fmt.Println("ssss--", buf)

	b := make([]byte, 100)
	ff, _ := os.Open("main.go")
	n, _ := ff.Read(b)
	fmt.Println(n)
	fmt.Println(string(b[:n]))

	//第二个参数，是指，从第几位开始读取
	n, _ = ff.ReadAt(b, 20)
	fmt.Println(n)
	fmt.Println(string(b[:n]))

	//获取文件夹下文件的列表
	dirs, err := os.Open("md")
	if err != nil {
		fmt.Println(err)
	}
	defer dirs.Close()
	//参数小于或等去0，表示读取所有的文件
	//另外一个只读取文件名的函数
	//fs, err := dirs.Readdirname(0)
	fs, err := dirs.Readdir(-1)
	if err == nil {
		for _, file := range fs {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Println("Readdir:", err)
	}

	//func (f *File) WriteString(s string) (ret int, err error)
	//写入字符串函数原型，哪个个函数比较快呢？？

	//p, _ := os.FindProcess(628)
	//fmt.Println(p)
	//p.Kill()
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	//参数也可以这么写 `c:\windows\system32\notepad.EXE`  用的是反单引号
	p, err := os.StartProcess("c:\\windows\\system32\\notepad.EXE", []string{"c:\\windows\\system32\\notepad.EXE", "d:/1.txt"}, attr)
	p.Release()
	time.Sleep(1000000000)
	p.Signal(os.Kill)
	os.Exit(10)

}
