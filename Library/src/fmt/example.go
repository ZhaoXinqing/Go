package main

import (
	"bytes"
	"fmt"
)

// fmt包实现了类似C语言printf和scanf的格式化I/O。格式化动作（'verb'）源自C语言但更简单
func main() {

	var buf bytes.Buffer
	var str string

	// 采用默认格式将其参数格式化并写入标准输出
	// 总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符
	// 返回写入的字节数和遇到的任何错误
	fmt.Println("打印内容", "Hello World!")

	// 采用默认格式将其参数格式化并写入标准输出
	// 如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格
	// 返回写入的字节数和遇到的任何错误
	fmt.Print("打印内容", "Hello World!", "\n")

	// 根据format参数生成格式化的字符串并写入标准输出
	// 返回写入的字节数和遇到的任何错误
	/*
		%b  二进制
		%o  八进制
		%d  十进制
		%x	十六进制
		%f	浮点数 3.141593
		%g  浮点数 3.141592653589793
		%e	浮点数 3.141593e+00
		%t 	布尔值
		%c  字符(rune)
		%s  字符串
		%q  带双引号的字符串"abc"或带单引号的'c'
		%v  变量的自然形式
		%T  变量的类型
		%%  字面上的%号标志
	 */
	fmt.Printf("打印内容%s", "Hello World!\n")

	// 采用默认格式将其参数格式化并写入w
	// 总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符
	// 返回写入的字节数和遇到的任何错误
	fmt.Fprintln(&buf, "Hello", "World!")

	// 采用默认格式将其参数格式化并写入w
	// 如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格
	// 返回写入的字节数和遇到的任何错误
	fmt.Fprint(&buf, "Hello", "World!")

	// 根据format参数生成格式化的字符串并写入w
	// 返回写入的字节数和遇到的任何错误
	fmt.Fprintf(&buf, "Hello %s", "World!")

	// 采用默认格式将其参数格式化，串联所有输出生成 并 返回一个字符串
	// 总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符
	fmt.Sprintln("Hello", "World!")

	// 采用默认格式将其参数格式化，串联所有输出生成 并 返回一个字符串
	// 如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格
	fmt.Sprint("Hello", "World!")

	// 根据format参数生成格式化的字符串 并 返回该字符串
	fmt.Sprintf("Hello %s", "World!")

	// 根据format参数生成格式化字符串并返回一个包含该字符串的错误
	fmt.Errorf("%s", "Error")

	// 从标准输入扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 换行视为空白。返回成功扫描的条目个数和遇到的任何错误
	// 如果读取的条目比提供的参数少，会返回一个错误报告原因
	fmt.Scan("Hello", "World!")

	// 类似Scan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置
	fmt.Scanln("Hello", "World!")

	// 从标准输入扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 返回成功扫描的条目个数和遇到的任何错误。
	fmt.Scanf("%s", "Hello")

	// 从r扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 换行视为空白。返回成功扫描的条目个数和遇到的任何错误
	// 如果读取的条目比提供的参数少，会返回一个错误报告原因
	fmt.Fscan(&buf, "Hello", "World!")

	// 类似Fscan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置
	fmt.Fscanln(&buf, "Hello", "World!")

	// 从r扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 返回成功扫描的条目个数和遇到的任何错误
	fmt.Fscanf(&buf, "%s", "Hello")

	// 从字符串str扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 换行视为空白。返回成功扫描的条目个数和遇到的任何错误
	// 如果读取的条目比提供的参数少，会返回一个错误报告原因
	fmt.Sscan(str, "Hello", "World!")

	// 类似Sscan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置
	fmt.Sscanln(str, "Hello", "World!")

	// 从字符串str扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数
	// 返回成功扫描的条目个数和遇到的任何错误
	fmt.Sscanf(str, "%s", "Hello")
}

// 向外输出（2、格式化占位符；3、获取输入）
// Print:
// 		- Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，
// 		- Printf函数支持格式化输出字符串，
// 		- Println函数会在输出内容的结尾添加一个换行符。

// Fprint:
// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
// 向标准输出写入内容
fmt.Fprintln(os.Stdout, "向标准输出写入内容")
fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    fmt.Println("打开文件出错，err:", err)
    return
}
name := "枯藤"
// 向打开的文件句柄中写入内容
fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
// 注意，只要满足io.Writer接口的类型都支持写入。

// Sprint:
// 		Sprint系列函数会把传入的数据生成并返回一个字符串。

// 获取输入：
// Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
// 可以在程序运行过程中从标准输入获取用户的输入。
// fmt.Scan，函数定签名如下：
func Scan(a ...interface{})(n int, err error)
// fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
// fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，
// fmt.Scanf为输入数据指定了具体的输入内容格式，
// 只有按照格式输入数据才会被扫描并存入对应变量:
func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
// Scanln类似Scan，
// 它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。

// bufio.NewReader
// 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，
// 这种情况下可以使用bufio包来实现。示例代码如下：
func bufioDemo() {
    reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    fmt.Print("请输入内容：")
    text, _ := reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
    fmt.Printf("%#v\n", text)
}

// Fscan系列
// 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
// 只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

// Sscan系列
// 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
// 只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)