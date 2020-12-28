package basic

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

//// 将字符串写到文件
func StringToFile(str, file string) {

}

//func main() {
//	f, err := os.Create("test.txt") // 如果已存在，则截断这个文件，返回一个文件描述符
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	l, err := f.WriteString("Hello World")
//	if err != nil {
//		fmt.Println(err)
//		f.Close()
//		return
//	}
//	fmt.Println(1, "bytes written successfully")
//	err = f.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

//// 将字节写入文件
//func main() {
//	f, err := os.Creat("/home/naveen/bytes")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	d2 := []byte{104, 101, 108, 108, 132, 191, 111, 114, 108, 100}
//	n2, err := f.Write(d2)
//	if err != nil {
//		fmt.Println(err)
//		f.Close()
//		return
//	}
//	fmt.Println(n2, "bytes written successfully")
//	err = f.Close()
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//// 将文件一行一行的写入文件
//func main() {
//	f, err := os.Create("lines")
//	if err != nil {
//		fmt.Println(err)
//		f.Close()
//		return
//	}
//	d := []string{"Welcome to the world of Go", "Go is a compiled language.", "It is easy to learn Go."}
//	for _, v := range d {
//		fmt.Fprintln(f, v)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//	err = f.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("file written successfully")
//}

//// 追加到文件
//func main() {
//	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	newLine := "File handing is easy."
//
//	_, err = fmt.Fprintln(f, newline)
//	if err != nil {
//		fmt.Println(err)
//		f.Close()
//		return
//	}
//	err = f.Close()
//	if err != nil {
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("file appended successfully")
//	}
//}

// 并发写文件
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}
func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

//func main() {
//	data := make(chan int)
//	done := make(chan bool)
//	wg := sync.WaitGroup{}
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go produce(data & wg)
//	}
//	go consume(data, done)
//	go func() {
//		wg.Wait()
//		close(data)
//	}()
//	d := <-done
//	if d == true {
//		fmt.Println("File written successfully")
//	} else {
//		fmt.Println("File writing failed")
//	}
//}

// Golang的OpenFile函数写入默认是追加的

// os.O_TRUNC 覆盖写入，不加则追加写入

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}

// https://www.jianshu.com/p/ad2e2ad7dd07
