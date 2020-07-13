import "container/list"

// list 包实现了一个双向链表，遍历一个链表
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}

// 通常函数会返回一个 error 值，调用它的代码应该判断这个错误是否等于 `nil`,来进行错误处理
// error 为 nil 时表示成功，非 nil 的 error 表示错误
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't conver number: %v", err)
}
fmt.Println("Converted integer number: %v", i)

// Go 的基本类型
bool, string

int int8 int16 int32(rune, "一个Unicode码") int64
uint uint8(byte) uint16 uint64 uintptr

float32 float64, complex64 complex128


unc main() {
	var i int 
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n",i, f, b, s)
}
// 0 0 false ""

func main() {
	var x,y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x,y,z)
}
// 3 4 5

func main() {
	v := 43
	fmt.Printf("v is of type %T\n",v)
}
// v is of type int

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello",World)
	fmt.Println("Happy",Pi,"Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
/*
Hello 世界
Happy 3.14 Day
Go rules? true
*/

const (
	Big = 1
	Small = Big
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x* 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
/*
11
0.1
0.1
*/


func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
// Go runs on windows.

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// When's Saturday?
// Too far away.

/*
defer 语句会延迟函数的执行直到上层函数返回。
延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
*/
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
// hello
// world

/*
defer 栈
延迟的函数调用被压入一个栈中。当函数返回时， 
会按照后进先出的顺序调用被延迟的函数调用。
*/
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i ++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
// {1 2}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
// 4

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
// {1 2} &{1 2} {1 0} {0 0}

// 类型 [n]T 是一个有 n 个类型为 T 的值的数组
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
/*
Hello World
[Hello World]
*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

/*
p == [2 3 5 7 11 13]
p[0] == 2
p[1] == 3
p[2] == 5
p[3] == 7
p[4] == 11
p[5] == 13
*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}
/*
p == [2 3 5 7 11 13]
p[1:4] == [3 5 7]
p[:3] == [2 3 5]
p[4:] == [11 13]
*/

//slice 的零值是 `nil`。一个 nil 的 slice 的长度和容量是 0。
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
*/

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// goroutine 是由 Go 运行时环境管理的轻量级线程。
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
ch <- v    // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。
（“箭头”就是数据流的方向。）
和 map 与 slice 一样，channel 使用前必须创建：ch := make(chan int)
默认情况下，在另一端准备好之前，发送和接收都会阻塞。
这使得goroutine可以在没有明确的锁或竞态变量的情况下进行同步。
*/
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
// -5 17 12

/*
channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
ch := make(chan int, 100)
向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞
*/
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
// 1
// 2

/*
发送者可以 close 一个 channel 来表示再没有值会被发送了。
接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：
当没有值可以接收并且 channel 已经被关闭，那么经过 v, ok := <-ch
之后 ok 会被设置为 `false`。
循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Print(i, "、")
	}
}

/*
select 语句使得一个 goroutine 在多个通讯操作上等待。
select 会阻塞，直到条件分支中的某个可以继续执行，
这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
// 0112358132134quit


// 当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
func main() {
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(50 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("    .")
			time.Sleep(10 + time.Millisecond)
		}
	}
}

/*
    .
    .
    .
tick
    .
    .
    .
    .
    .
tick
    .
    .
    .
    .
tick
    .
    .
    .
    .
    .
tick
    .
    .
    .
    .
BOOM
*/

package main

import "code.google.com/p/go-tour/tree"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int)

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool

func main() {
}

// 使用 Go 的并发特性来并行执行 web 爬虫。
//修改 Crawl 函数来并行的抓取 URLs，并且保证不重复。
package main
import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string,depth int,fectcher Fetcher) {
	if depth <= 0 {
		return
	}
	body,urls,err := fectcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found:%s %q\n",url, body)
	for _,u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}


import {
	"fmt"
	"os"
}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

for {
	// ... 无限循环
}

line := input.Text()
counts[line] = counts[line] + 1

input := bufio.NewScanner(os.Stdin)

%d		int
%x %o %b 16 8 2 int
%f %g %e 浮点，Π，e
%t 布尔
%c rune
%s string
%v 将任意变量以易读的形式打印出来
%T 打印变量的类型

Printf默认不会在输出内容后加上换行符。
按照惯例，用来格式化的函数都会在末尾以f字母结尾（译注：f后缀对应format或fmt缩写），
比如log.Printf，fmt.Errorf，同时还有一系列对应以ln结尾的函数（译注：ln后缀对应line缩写），
这些函数默认以%v来格式化他们的参数，并且会在输出结束后在最后自动加上一个换行符。

package main

// 从标准输入得到一些文件名，然后用os.Open函数来打开每一个文件获取内容。
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) = 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, arg := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n, line)
		}
	}
}

func countlines(f *os.File,counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var

内建常量：true false iota nil

内建类型：int int8 int16 int32 int64
		uint uint8 uint16 uint32 uint64 uintptr
		float32 float64 complex128 complex64
		bool byte rune string error

内建函数：make len cap new append copy close delete
		complex real imag
		panic recover

		TCP客户端
		一个TCP客户端进行TCP通信的流程如下：
		
		建立与服务端的链接
		进行数据收发
		关闭链接

// TCP server 端
func process(conn net.Conn) {
	defer conn.Close()  // 关闭连接
	for {
		reader := buffio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read{buf[:]}  // 读取数据
		if err != nil {
			fmt.Println("连接客户端失败，错误信息："，err)
		}
		recvStr := string{buf[:n]}
		fmt.Println("收到客户端信息：", recvStr)
		conn.Write([]byte{recvStr})  //发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp","127.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败，错误：", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("建立连接失败，错误：", err)
			continue
		}
		go process(conn)  // 启动一个goroutine处理连接
	}
}

// tcp客户端
func main() {
	conn, err := net.Dial("tcp","127.0.0.0:8888")
	if err != nil {
		fmt.Println("连接失败，错误：", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdout)
	for {
		input,_ := inputReader.ReadString('\n')
		inputInfo := string.Trim(input,"\r\n")
		if string.ToUpper(inputInfo) == "q" {
			return // 如果输入q就退出
		}
		_, err = conn.Write([]byte(inputInfo))  // 发送数据
		if err := nil{
			return
		}
		buf := [512]byte{}
		n, err != nil{
			fmt.Println("接受失败，错误：", err)
			return
		}
		fmt.Println(astring(buf[:n]))
	}
}

UDP服务端
//服务端
func main()  {
    listen,err := net.ListenUDP("udp",&net.UDPAddr{
        IP:net.IPv4(0,0,0,0),
        Port:8888,
    })
    if err != nil{
        fmt.Println("监听失败，错误：",err)
        return
    }
    defer listen.Close()
    for {
        var data [1024]byte
        n,addr,err := listen.ReadFromUDP(data[:])
        if err != nil{
            fmt.Println("接收udp数据失败，错误：",err)
            continue
        }
        fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
        _ ,err = listen.WriteToUDP(data[:n],addr)   //发送数据
        if err != nil{
            fmt.Println("发送数据失败，错误：",err)
            continue
        }
    }
}
UDP客户端
//客户端
func main()  {
    socket,err := net.DialUDP("udp",nil,&net.UDPAddr{
        IP:net.IPv4(0,0,0,0),
        Port:8888,
    })
    if err != nil{
        fmt.Println("连接服务器失败，错误：",err)
        return
    }
    defer socket.Close()
    sendData := []byte("hello world!")
    _,err = socket.Write(sendData)
    if err != nil{
        fmt.Println("发送数据失败，错误：",err)
        return
    }
    data := make([]byte,4096)
    n,remoteAddr,err := socket.ReadFromUDP(data)
    if err != nil{
        fmt.Println("接受数据失败，错误：",err)
        return
    }
    fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}


// http server
func sayHi(w http.ResponseWriter,r *http.Request)  {
    fmt.Fprintln(w,"你好，ares！")
}
func main()  {
    http.HandleFunc("/",sayHi)
    err := http.ListenAndServe(":8888",nil)
    if err != nil{
        fmt.Println("Http 服务建立失败，err：",err)
        return
    }
}

// http server
func sayHi(w http.ResponseWriter,r *http.Request)  {
    fmt.Fprintln(w,"你好，ares！")
}
func main()  {
    http.HandleFunc("/",sayHi)
    err := http.ListenAndServe(":8888",nil)
    if err != nil{
        fmt.Println("Http 服务建立失败，err：",err)
        return
    }
}

func main() {
    resp, err := http.Get("https://www.baidu.com/")
    if err != nil {
        fmt.Println("get failed, err:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("%T\n",body)
    fmt.Println(string(body))
}

// Encode 将消息编码
func Encode(message string)([]byte ,error)  {
    // 读取消息的长度，转换成int32类型（占4个字节）
    var length = int32(len(message))
    var pkg = new(bytes.Buffer)
    //写入消息头
    err := binary.Write(pkg,binary.LittleEndian,length)
    if err != nil{
        return nil,err
    }
    //写入消息实体
    err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
    if err != nil{
        return nil,err
    }
    return pkg.Bytes(),nil
}

// Decode 消息解码
func Decode(reader *bufio.Reader)(string,error)  {
    //读取消息长度
    lengthByte,_ := reader.Peek(4) //读取前4个字节数据
    lengthBuff := bytes.NewBuffer(lengthByte)
    var length int32
    err := binary.Read(lengthBuff,binary.LittleEndian,&length)
    if err != nil{
        return "",err
    }
    // Buffered返回缓冲中现有的可读取的字节数。
    if int32(reader.Buffered()) < length+4{
        return "",err
    }
    //读取真正的消息数据
    pack := make([]byte,int(4+length))
    _,err = reader.Read(pack)
    if err != nil{
        return "",err
    }
    return string(pack[4:]),nil
}

// server端：
func process(conn net.Conn)  {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        msg,err := proto.Decode(reader)
        if err == io.EOF{
            return
        }
        if err != nil{
            fmt.Println("decode 失败,err",err)
            return
        }
        fmt.Println("收到client数据:",msg)
    }
}
func main()  {
    listen,err := net.Listen("tcp","127.0.0.1:8888")
    if err != nil{
        fmt.Println("监听失败,err",err)
        return
    }
    defer listen.Close()
    for {
        conn,err := listen.Accept()
        if err != nil{
            fmt.Println("接受失败,err",err)
            continue
        }
        go process(conn)
    }
}

// client端:
func main()  {
    conn,err := net.Dial("tcp","127.0.0.1:8888")
    if err != nil{
        fmt.Println("dial失败，err",err)
        return
    }
    defer conn.Close()
    for i:=0;i<20;i++{
        msg := "Hello Ares!"
        data,err := proto.Encode(msg)
        if err != nil{
            fmt.Println("encode失败，err",err)
            return
        }
        conn.Write(data)
    }
}
goroutine
goroutine
goroutine
import "unsafe"
fmt.Println(unsafe.Sizeof(float64(0))) //  "8"

类型 大小
bool 1个字节
intN, uintN, floatN, complexN N/8个字节(例如float64是8个字节)
int, uint, uintptr 1个机器字
*T 1个机器字
string 2个机器字(data,len)
[]T 3个机器字(data,len,cap)
map 1个机器字
func 1个机器字
chan 1个机器字
interface 2个机器字(type,value)

pb := (*int16)(unsafe.Pointer(
	uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
*pb = 42
fmt.Println(x.b) // "42"

reflect
reflect
reflect

DeepEqual
DeepEqual
DeepEqual
// reflect包的DeepEqual函数可以对两个值进行深度相等判断,
// 用DeepEqual函数比较两个字符串数组是否相等
func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c",":")
	want := []string{"a","b","c"};
	if !reflect.DeepEqual(got,want) { /* ... */}
}

var a, b []string = nil, []string{}
fmt.Println(reflect.DeepEqual(a,b)) // "false"

// Equal resports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x),reflect.ValueOf(y),seen)
}

type comparsion struct {
	x, y unsafe.Pointer
	treflect.Type
}

foreign-function interfaces

deflate
compress

package gzip // compress/gzip
func NewWriter(w io.Writer) io.WriteCloser
func NewWriter(r io.Reader) (io.ReaderCloser, error)

func TestName(t *testing.T) {
	// ...
}

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
		}
	}
	return true
}

package word
import "testing"
func TestPalindrome(t *testing.T) {
	if !IsPalindrome("datartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayka") {
		t.Error(`Ispalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrom") = true`)
	}
}

//切片扩容
func main () {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
// [0 0 0 0 0 1 2 3]
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt,Println(s)
}
// [1 2 3 4]

// new() 和 make() 的区别？
/* 
	1、new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
	2、new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话
说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
	3、make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过
	初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
*/

const (
	_ = iota
	c1 int = (10 * iota)
	c2
	d = iota
)
func main() {
	fmt.Printf("%d - %d - %d", c1,c2,d)
}
// 10 - 20 - 3

//并发，引用，输出10
const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i 
			mu.Unlock()
		}()
	}
		
}


type Student struct{
     Age int
 }

 func main(){
    s := Student{30}
    modify1(s)
	fmt.Println(s) //{30}
	
    modify2(&s)
    fmt.Println(s) //{32}
 }

 func modify1(s Student){
    s.Age = 31
 }

 func modify2(s *Student){
    s.Age = 32
 }
// 递归、阶乘
package main

import "fmt"

func Factorial(x int)(result int){
	if x == 0{
		result = 1
	} else {
		result =x * Factorial(x-1)
	}
	return
}

func main() {
	var i int = 15
	fmt.Println("%d 的阶乘是 %d\n", i, Factorial(i))
}

//递归、斐波那契
package main

import "fmt"

func fibonacc(n int) int {
	if n < 2 {
		return n
	}
	return fibonacc(n-2) + fibonacc(n-1)
}

func main() {
	var i int
	for i = 0;i < 10;i++ {
		fmt.Println( fibonacc(i))
	}
}

//关键字
func
struct

//全局变量
var str string
var str = ""

//局部变量的初始化
var i int = 10
var i = 10
i := 10

//访问p的成员变量name
p.name
（*p).name

//接口和类
"一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口"
"实现类的时候，只需要关系自己应该提供哪些方法，不用再纠结接口需要拆的多细才合理"
"接口由使用方法自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口"

//字符串连接
str := "abc" + "123"
fmt.Sprintf("abc%d",123)

//关于协程
"协程和线程都可以实现程序的并发执行"
"通过channel来进行协程间的通信"

//init 函数
"一个包中，可以包含多个init 函数"
"程序编译时，先执行导入包的init函数，再执行本包内的init函数"
"init 函数不可以被其它函数调用"

//循环语句
"for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环"
"for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量"

//函数定义
func add(arg ...int)int{
	sum := 0
	for _,arg := range args{
		sum += arg
	}
	return sum
}
"add函数调用"
add(1,2)
add(1,3,7)
add([]int{1,3,7}...)

//类型转换
type Mylnt int
var i int = 1
var jMylnt = Mylnt(i)


//const 常量定义的使用方法
const Pi float64 =3.14159265358979323846
const zero = 0.0

const(
	size int64 = 1024
	eof = -1
)

const u,v float32 = 0,3
const a,b,c = 3,4,"foo"

//布尔变量赋值
b = true
b = (1 ==2)

//求运行结果
func main(){
	if (true){
		defer fmt.Printf("1")
	}else{
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}
// 31

//switch语句
"单个case中，可以出现多个结果选项"
"只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case"

//golang中没有隐藏的this指针，含义是
"方法施加的对象显示传递，没有被隐藏起来"
"golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达"
"方法施加的对象不需要非得是指针，也不用非得叫this"

//golang中的引用类型
"数组切片、map、channel、interface"

//golang中的指针运算
"通过 "&" 取指针的地址"
"通过 "*" 取指针指向的数据"

//main函数
"不能带参数"
"不能定义返回值"
"所在的包必须是main包"
"可以使用flag包来获取和解析命令行参数"

//正确赋值
var x interface{} = nil
var x error = nil

//切片初始化
s := make([]int,0)
s := make([]int,05,10)
s := []int{1,2,3,4,5}

//切片中删除一个元素
func (s *Slice)Remove(value interface{})error{
	for i,v := range *s{
		if isEqual(value,v){
			*s = append((*s)[:i],(*s)[i +1:]...)
			return nil
		}
	}
	return ERR_ELEM_NT_NEXT
}

//局部变量整形切片x的赋值
x := []int{
	1,2,3,
	4,5,6,
}

x := []int{
	1,2,3,
	4,5,6}

x := []int{1,2,3,4,5,6,}

//变量的自增自减操作
i := 1
i ++

i := 1
i --

//函数声明
func f(a,b int)(value int,err error)
func f(a int,b int)(value int,err error)
func f(a int,b int)(int, int, error)

//add函数调用代码
func main(){
	var a lnteger = 1
	var b lnteger = 2
	var i interface{} = &a
	sum := i.(*lnteger).Add(b)
	fmt.Printf(sum)
}
//"Add定义函数"
type lnteger int
func (a lnteger) Add (b lnteger)lnteger{
	return a + b
}

type lnteger int
func (a *lnteger) Add(b lnteger) integer{
	return *a + b
}

//调用函数为
func main(){
	var a lnteger = 1
	var b lnteger = 2
	var i interface{} = a
	sum := i.(lnteger).Add(b)
	fmt.Println(sum)
}
//add定义函数为
type lnteger int
func (a lnteger)Add(b lnteger)lnteger{
	return a + b
}

// 关于GetPodAction定义，
type Fragment interface{
	Exec(translnfo *Translnfo)error
}
type GetPodAction struct{
}
func( g GetPodAction)Exec(translnfo *Translnfo)error{
	...
	return nil
}
//下面赋值正确的是
var fragment Fragment = new(GetPodAction)
var fragment Fragment = &GetPodAction{}
var fragment Fragment = GetPodAction{}

// 关于GoMock
"可以对interface打桩"
"打桩后的依赖注入可以通过GoStub完成"

//关于接口
"只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值"
"如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A"
"接口查询是否成功，要在运行期才能够确定"

//channel语法
var ch chan int
ch := make(chan int)
<- ch

//同步锁
"当一个goroutine获得Mutex后，其它goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex"
"RWMutex在读锁占用的情况下，会阻止写，但不阻止读"
"RWMutex在写锁占用的情况下，会阻止任何其它goroutine（无论读和写）进来，整个锁相当于由该goroutine独占"

//golang中数据类型不能转化为有效的JSON文本的是：
"channel complex 函数"

//go vendor
"基本思路是将引用的外部包的源码放在当前工程的vendor目录下"
"编译go代码会优先从vendor目录先寻找依赖包"
"有了vendor目录后，打包当前的工程代码到其它机器的$GOPATH/src下都可以通过编译"

//falg是bool型变量，if表达式符合编码规范的是
if falg
if ！flag

//value是整型变量，if表达式符合编码规范的是
if value == 0
if value ！= 0

//函数返回值错误设计
"失败原因只有一个，则返回bool"
"失败原因超过一个，则返回error"
"如果没有失败原因，则不超过bool或error"
"如果重试几次可以避免失败，则不要立即返回bool或error"

//异常设计
"在开发阶段，坚持速错，让程序异常崩溃"
"在程序部署后，应恢复异常避免程序终止"
"对于不应该出现的分支，使用异常处理"

// slice 和 map 操作
var s []int
s = append(s,1)

var s []int
s = make([]int,0)
s = append(s,1)

var m map[string]int
m = make(map[string]int)
m["one"] = 1

//channel特性
"给一个nil channel 发送数据，造成永远阻塞"
"从一个nil channel 接受数据，造成永远阻塞"
"给一个已经关闭的channel发送数据，引起panic"
"从一个已经关闭的channel接收数据，如果缓冲区中为空，则返回一个零值"

//channel 缓冲、冲突
"无缓冲的channel是同步的，而有缓冲的channel是非同步的"

//异常的触发
"空指针解析"
"下标越界"
"除数为零"
"调用panic函数"

//cap函数的适用类型
"array、slice、channel"

//beego框架
"是一个golang实现的轻量级HTTP框架"
"可以通过注释路由、正则路由等多种方式完成url路由注入"
"可以使用bee new工具生成空工程，然后 使用bee run命令自动热编译"

//goconvey
"是一个支持golang的单元测试框架"
"能够自动监控文件并启动测试，并可以将测试结果实时输出到web界面"
"提供丰富的断言简化测试用例的编写"

//go vet
"是golang自带工具go tool vet的封装"
"可以使用绝对路径、相对路径或相对GOPATH的路径指定待检测的包"
"可以检测出死代码"

//map
"map反序列化时json.unmarshal的入参必须为map的地址"
"在函数调用中传递map，则子函数中对map元素的增加会导致父函数中map的修改"
"在函数调用中传递map，则子函数中对map元素的修改会导致父函数中map的修改"
"能使用内置函数delete删除map的元素"

//GoStub
"可以对全局变量打桩"
"可以对函数打桩"
"不可以对类的成员方法打桩"
"可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为"

//select机制
"用来处理异步IO问题"
"最大的一条限制就是每个case语句里必须是一个IO操作"
"golang在语言级别支持select关键字"

//内存泄漏
"golang中检测内存泄露主要依靠的是pprof包"
"应定期使用浏览器来查看系统的实时内存信息，及时发现内存泄露问题"


填空题

//声明一个整型变量i
var i int

//声明一个含有10个元素的整型数组a
var a [10]int

//声明一个整型组切片s
var s []int

//声明一个整型指针变量p
var p *int

//声明一个key为字符串value的整型的map变量m
var m map[string]int

//声明一个入参和返回值均为整型的函数变量f
var f func(a int)int

//声明一个只用于读取int数据的单向channel变量ch
var ch <-chan int 

//假设源文件的命名为slice.go,则测试文件的命名为
slice_test.go

//go test要求测试函数的前缀必须命名为
Test

//求程序运行结果
for i := 0; i < 5; i++{
	defer fmt.Printf("%d",i)
}
//4 3 2 1 0

func main(){
	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Println(x)
}
//21

func main(){
	strs := []string{"one","two","three"}
	for _,s := range strs{
		go func(){
			time.Sleep(1 *time.Second)
			fmt.Printf("%s",s)
		}()
	}
	time.Sleep(3 *time.Second)
}
//three threethree

func main(){
	x := []string{"a","b","c"}
	for v := range x{
		fmt.Print(v)
	}
}
//012

func main(){
	x := []string{"a","b","c"}
	for _,v := range x {
		fmt.Print(v)
	}
}
//abc

func main(){
	i := 1
	j := 2
	i,j = j,i
	fmt.Printf("%d%d\n",i,j)
}
//21

func incr(p *int)int{
	*p++
	return *P
}
func main(){
	v := 1
	incr(&v)
	fmt.Println(v)
}
//2

//启动一个goroutine的关键字
go

//求运行结果
type Slice []int
func NewSlice() Slice{
	return make(Slice,0)
}
func (s *Slice) Add(elem int) *Slice{
	*s = append(*s,elem)
	fmt.Print(elem)
	return s
}
func main(){
	s := NewSlice()
	defer s.Add(3)
}
//132


判断对错
//数组是一个值类型
//使用map不需要引入任何库
//内置函数delete不能删除数组切片内的元素
//指针不是基础类型
//interface{}是可以指向任意对象的Any类型

//下面文件操作的代码可能触发异常
file,err := os.Open("test.go")
defer file.Close()
if err != nil{
	fmt.Println("open file failed:",err)
	return
}

//Golang支持自动垃圾回收
//Golang支持反射，反射最常见的使用场景是做对象的序列化

//Golang可以复用C/C++的模块，这个功能叫Cgo（错误）
//通过成员变量或函数首字母的大小写来决定其作用域

//下面的程序的运行结果不是xello
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}

//golang支持goto语句

//面代码中的指针p为野指针，因为返回的栈内存在函数结束时会被释放（错误）
type TimesMatcher struct {
	base int
}

func NewTimesMatcher(base int) *TimesMatcher{
	return &TimesMatcher{base:base}
}

func main() {
	p := NewTimesMatcher(3)
	...
}

//匿名函数可以直接赋值给一个变量或者直接执行

//如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
//---可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量

//在函数的多返回值中，如果有error或bool类型，则一般放在最后一个
//错误是业务过程的一部分，而异常不是
//函数执行时，如果由于panic导致了异常，则延迟函数会执行

//同级文件的包名不允许有多个
//golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承
//使用for range迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的
//switch后面可以不跟表达式


//结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
//------因此在decode时这些非导出变量的值为其类型的零值

// golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名

//当函数deferDemo返回失败时，能destroy已create成功的资源（）

func deferDemo() error {
	err := createResource1()	
	if err != nil {
	    return ERR_CREATE_RESOURCE1_FAILED
	}
	defer func() {
		if err != nil {
			destroyResource1()
		}
	}()
	
	err = createResource2()
	if err != nil {
		return ERR_CREATE_RESOURCE2_FAILED
	}
	
	defer func() {
		if err != nil {
			destroyResource2()
		}
	}()
	
	err = createResource3()
	if err != nil {
		return ERR_CREATE_RESOURCE3_FAILED
	}
	return nil
}


最近在很多地方看到了golang的面试题，看到了很多人对Golang的面试题心存恐惧，也是为了复习基础，我把解题的过程总结下来。

面试题

1 写出下面代码输出内容。

package main
import (   
"fmt"
)
funcmain() {
    defer_call()
}
funcdefer_call() {
    deferfunc() {fmt.Println("打印前")}()
    deferfunc() {fmt.Println("打印中")}()
    deferfunc() {fmt.Println("打印后")}()
    panic("触发异常")
}

考点：defer执行顺序
解答：
defer 是后进先出。
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

打印后
打印中
打印前
panic: 触发异常

2 以下代码有什么问题，说明原因。

type student struct {
    Name string
    Age  int
}
funcpase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou",Age: 24},
        {Name: "li",Age: 23},
        {Name: "wang",Age: 22},
    }    for _,stu := range stus {
        m[stu.Name] =&stu
    }
}

考点：foreach
解答：
这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。 就像想修改切片元素的属性：

for _, stu := rangestus {
    stu.Age = stu.Age+10}

也是不可行的。 大家可以试试打印出来：

func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou",Age: 24},
        {Name: "li",Age: 23},
        {Name: "wang",Age: 22},
    }    
    // 错误写法
    for _,stu := range stus {
        m[stu.Name] =&stu
    }    
     fork,v:=range m{        
      println(k,"=>",v.Name)
    }    
      // 正确
    for i:=0;i<len(stus);i++ {
       m[stus[i].Name] = &stus[i]
    }    
     fork,v:=range m{        
       println(k,"=>",v.Name)
    }
}

3 下面的代码会输出什么，并说明原因

func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}
    wg.Add(20)   for i := 0; i < 10; i++ {        
         gofunc() {
           fmt.Println("A: ", i)
           wg.Done()
        }()
    }    
        for i:= 0; i < 10; i++ {        
           gofunc(i int) {
           fmt.Println("B: ", i)
           wg.Done()
        }(i)
    }
    wg.Wait()
}

考点：go执行的随机性和闭包
解答：
谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。 但是A:均为输出10，B:从0~9输出(顺序不定)。 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
4 下面代码会输出什么？

type People struct{}func (p *People)ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func(p*People)ShowB() {
    fmt.Println("showB")
}
typeTeacher struct {
    People
}
func(t*Teacher)ShowB() {
    fmt.Println("teachershowB")
}
funcmain() {
    t := Teacher{}
    t.ShowA()
}

考点：go的组合继承
解答：
这是Golang的组合模式，可以实现OOP的继承。 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
showAshowB

5 下面代码会触发异常吗？请详细说明

func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chanint, 1)
    string_chan := make(chanstring, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {   
            case value := <-int_chan:
       fmt.Println(value)
          casevalue := <-string_chan:        
          panic(value)
    }
}

考点：select随机性
解答：
select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

select 中只要有一个case能return，则立刻执行。
当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
如果没有一个case能return则可以执行”default”块。

6 下面代码输出什么？

funccalc(indexstring, a, bint) int {
    ret := a+ b
    fmt.Println(index,a, b, ret)
    return ret
}
funcmain() {   
      a := 1
    b := 2
    defer calc("1", a,calc("10", a, b))    a = 0
    defer calc("2", a,calc("20", a, b))    b = 1
}

考点：defer执行顺序
解答：
这道题类似第1题 需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用
calc("10",1,2)==>10,1,2,3 执行index:2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2 执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2最后执行index:1==>calc("1",1,3)==>1,1,3,4

10 1 2 320 0 2 22 0 2 21 1 3 4

7 请写出以下输入内容

funcmain() {    
       s := make([]int,5)
    s = append(s,1, 2, 3)
    fmt.Println(s)
}

考点：make默认值和append
解答：
make初始化是由默认值的哦，此处默认值为0

[00000123]

大家试试改为:

s := make([]int, 0)
s = append(s, 1, 2, 3)
fmt.Println(s)//[1 2 3]

8 下面的代码有什么问题?

type UserAges struct {
    ages map[string]int
    sync.Mutex
}
func(ua*UserAges)Add(name string, age int) {
    ua.Lock()  
       deferua.Unlock()
    ua.ages[name] = age
}
func(ua*UserAges)Get(name string)int {    
      ifage, ok := ua.ages[name]; ok {        
         return age
    }  
      return-1
}

考点：map线程安全
解答：
可能会出现

fatal error: concurrent mapreadandmapwrite.

修改一下看看效果

func (ua *UserAges)Get(namestring)int {
    ua.Lock()    
     deferua.Unlock()    
     ifage, ok := ua.ages[name]; ok {        
          return age
    }    
       return-1
}

9.   下面的迭代会有什么问题？

func (set *threadSafeSet)Iter()<-chaninterface{} {
    ch := make(chaninterface{})   
              gofunc() {
        set.RLock()   
            for elem := range set.s {
           ch <- elem
        }     
             close(ch)
        set.RUnlock()
    }()
     return ch
}

考点：chan缓存池
解答：
看到这道题，我也在猜想出题者的意图在哪里。 chan?sync.RWMutex?go?chan缓存池?迭代? 所以只能再读一次题目，就从迭代入手看看。 既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。 我们把代码恢复为可以运行的方式，看看效果

package main
import (   
      "sync"
    "fmt")//下面的迭代会有什么问题？type threadSafeSet struct {
    sync.RWMutex
    s []interface{}
}
func(set*threadSafeSet)Iter() <-chaninterface{} {    
//ch := make(chan interface{}) // 解除注释看看！
    ch := make(chaninterface{},len(set.s))   
gofunc() {
        set.RLock()       
forelem,value := range set.s {
           ch <- elem            
println("Iter:",elem,value)
        }       close(ch)
        set.RUnlock()
    }()    
return ch
}
funcmain() {
    th:=threadSafeSet{
        s:[]interface{}{"1","2"},
    }
    v:=<-th.Iter()
    fmt.Sprintf("%s%v","ch",v)
}

10 以下代码能编译过去吗？为什么？

package main
import (   "fmt")
typePeople interface {
    Speak(string) string
}
typeStduent struct{}
func(stu*Stduent)Speak(think string)(talk string) {    
ifthink == "bitch" {
        talk = "Youare a good boy"
    } else {
        talk = "hi"
    }
    return
}
funcmain() {
    var peoPeople = Stduent{}
    think := "bitch"
   fmt.Println(peo.Speak(think))
}

考点：golang的方法集
解答：
编译不通过！ 做错了！？说明你对golang的方法集还有一些疑问。 一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。

11 以下代码打印出来什么内容，说出为什么。

package main
import (   "fmt")
typePeople interface {
    Show()
}
typeStudent struct{}
func(stu*Student)Show() {
}
funclive()People {
    var stu*Student
    return stu
}
funcmain() {   if live() == nil
{
        fmt.Println("AAAAAAA")
    } else {
        fmt.Println("BBBBBBB")
    }
}

考点：interface内部结构
解答：
很经典的题！ 这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

varininterface{}

另一种如题目：

type People interface {
    Show()
}

他们的底层结构如下：

type eface struct {      //空接口
    _type *_type        //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)}
typeiface struct {      //带有方法的接口
    tab  *itab          //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)}
type_type struct {
    size       uintptr //类型大小
    ptrdata    uintptr //前缀持有所有指针的内存大小
    hash       uint32  //数据hash值
    tflag     tflag
    align      uint8   //对齐
    fieldalign uint8   //嵌入结构体时的对齐
    kind       uint8   //kind 有些枚举值kind等于0是无效的
    alg       *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte   str       nameOff
    ptrToThis typeOff
}type itab struct {
    inter  *interfacetype //接口类型
    _type  *_type         //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr     //可变大小方法集合}

可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil， 所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口 结果：

BBBBBBB
12.是否可以编译通过？如果通过，输出什么？

func main() {
    i := GetValue() switch i.(type) { 
        caseint:       
        println("int")   
        casestring:       
        println("string")   
        caseinterface{}:       
        println("interface")   
        default:       
         println("unknown")
    }
}
funcGetValue()int {   
return1
}

解析
考点：type

编译失败，因为type只能使用在interface

13.下面函数有什么问题？

func funcMui(x,y int)(sum int,error){    
returnx+y,nil
}

解析
考点：函数返回值命名
在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。 如果返回值有有多个返回值必须加上括号； 如果只有一个返回值并且有命名也需要加上括号； 此处函数第一个返回值有sum名称，第二个未命名，所以错误。

14.是否可以编译通过？如果通过，输出什么？

package mainfunc main() {    println(DeferFunc1(1)) println(DeferFunc2(1)) println(DeferFunc3(1))
}func DeferFunc1(i int)(t int) {
    t = i   deferfunc() {
        t += 3
    }() return t
}
funcDeferFunc2(i int)int {
    t := i  deferfunc() {
        t += 3
    }() return t
}
funcDeferFunc3(i int)(t int) {   deferfunc() {
        t += i
    }() return2}

解析
考点:defer和函数返回值
需要明确一点是defer需要在函数结束前执行。 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数 DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4; DeferFunc2函数中t的作用域为函数，返回1;DeferFunc3返回3

15.是否可以编译通过？如果通过，输出什么？

funcmain() {    list := new([]int)
    list = append(list,1)
    fmt.Println(list)
}

解析
考点：new

list:=make([]int,0)

16.是否可以编译通过？如果通过，输出什么？

package mainimport "fmt"funcmain() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s1 = append(s1,s2)
    fmt.Println(s1)
}

解析
考点：append
append切片时候别漏了'…'

17.是否可以编译通过？如果通过，输出什么？

func main() {
    sn1 := struct {
        age  int
        name string
    }{age: 11,name: "qq"}
    sn2 := struct {
        age  int
        name string
    }{age: 11,name: "qq"}  if sn1== sn2 {
        fmt.Println("sn1== sn2")
    }
    sm1 := struct {
        age int
        m   map[string]string
    }{age: 11, m:map[string]string{"a": "1"}}
    sm2 := struct {
        age int
        m   map[string]string
    }{age: 11, m:map[string]string{"a": "1"}} 
           if sm1 == sm2 {
        fmt.Println("sm1== sm2")
    }
}

解析
考点:结构体比较
进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。

sn3:= struct {
    name string
    age  int
}
{age:11,name:"qq"}

sn3与sn1就不是相同的结构体了，不能比较。 还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。

可以使用reflect.DeepEqual进行比较

if reflect.DeepEqual(sn1, sm) {
    fmt.Println("sn1==sm")
}else {
    fmt.Println("sn1!=sm")
}

所以编译不通过： invalid operation: sm1 == sm2

18.是否可以编译通过？如果通过，输出什么？

func Foo(x interface{}) {    if x== nil {
        fmt.Println("emptyinterface")      
          return
    }
    fmt.Println("non-emptyinterface")
}
       funcmain() {   
       var x *int = nil
    Foo(x)
}

解析
考点：interface内部结构

non-emptyinterface


19.是否可以编译通过？如果通过，输出什么？

func GetValue(m map[int]string, id int)(string, bool) {    
         if _,exist := m[id]; exist {        
           return"存在数据", true
    }  
         returnnil, false}funcmain() {
    intmap:=map[int]string{   
1:"a",       
2:"bb",       
3:"ccc",
    }
    v,err:=GetValue(intmap,3)
    fmt.Println(v,err)
}

解析
考点：函数返回值类型
nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。报:cannot use nil as type string in return argument.

20.是否可以编译通过？如果通过，输出什么？

const (
    x = iota
    y
    z = "zz"
    k
    p = iota)
funcmain() 
{
    fmt.Println(x,y,z,k,p)
}

解析
考点：iota
结果:

0 1 zz zz 4

21.编译执行下面代码会出现什么?

package mainvar(
    size :=1024
    max_size = size*2)
funcmain() {    
println(size,max_size)
}

解析
考点:变量简短模式
变量简短模式限制：

定义变量同时显式初始化
不能提供数据类型
只能在函数内部使用
结果：

syntaxerror: unexpected :=

22.下面函数有什么问题？

package main
const cl = 100
var bl   = 123
funcmain() {    
println(&bl,bl)   
println(&cl,cl)
}

解析
考点:常量
常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，

cannot take the address of cl

23.编译执行下面代码会出现什么?

package main
funcmain() {    
for i:=0;i<10;i++  {
    loop:       
println(i)
    }    gotoloop
}

解析
考点：goto
goto不能跳转到其他函数或者内层代码

goto loop jumps intoblock starting at

24.编译执行下面代码会出现什么?

package main
import"fmt"
funcmain() {    
 typeMyInt1 int    
 typeMyInt2 = int
    var i int =9
    var i1MyInt1 = i
    var i2MyInt2 = i
    fmt.Println(i1,i2)
}

解析
考点：**Go 1.9 新特性 Type Alias **
基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。 MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转； MyInt2称之为alias，可以直接赋值。

结果:

cannot use i (typeint) astype MyInt1 in assignment

25.编译执行下面代码会出现什么?

package main
import"fmt"
typeUser struct {
}
typeMyUser1 User
typeMyUser2 = User
func(iMyUser1)m1(){
    fmt.Println("MyUser1.m1") 
}
func(iUser)m2(){
    fmt.Println("User.m2")
}
funcmain() {
    var i1MyUser1
    var i2MyUser2
    i1.m1()
    i2.m2()
}

解析
考点：**Go 1.9 新特性 Type Alias **
因为MyUser2完全等价于User，所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。 但是

i1.m2()

是不能执行的，因为MyUser1没有定义该方法。 结果:

MyUser1.m1User.m2

26.编译执行下面代码会出现什么?

package main
import"fmt"
type T1 struct {
}
func(tT1)m1(){
    fmt.Println("T1.m1")
}
type T2= T1
typeMyStruct struct {
    T1
    T2
}
funcmain() {
    my:=MyStruct{}
    my.m1()
}

解析
考点：**Go 1.9 新特性 Type Alias **
是不能正常编译的,异常：

ambiguousselectormy.m1

结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。 改为:

my.T1.m1()
my.T2.m1()

type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，别名类型也怎么用，保留源类型的所有方法、字段等。

27.编译执行下面代码会出现什么?

package main
import (   
       "errors"
    "fmt")
varErrDidNotWork = errors.New("did not work")
funcDoTheThing(reallyDoItbool)(errerror) {    
ifreallyDoIt {
        result, err:= tryTheThing()        
if err!= nil || result != "it worked" {
           err = ErrDidNotWork
        }
    }    return err
}
functryTheThing()(string,error) {    
return"",ErrDidNotWork
}
funcmain() {
    fmt.Println(DoTheThing(true))
    fmt.Println(DoTheThing(false))
}

解析
考点：变量作用域
因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：

改为：

func DoTheThing(reallyDoIt bool)(errerror) {    
varresult string
    ifreallyDoIt {
        result, err =tryTheThing()        
if err!= nil || result != "it worked" {
           err = ErrDidNotWork
        }
    }    return err
}

28.编译执行下面代码会出现什么?

package main
functest() []func() {    
varfuns []func()
    fori:=0;i<2;i++  {
        funs = append(funs,func() {           
           println(&i,i)
        })
    }    returnfuns
}
funcmain(){
    funs:=test()    
       for_,f:=range funs{
        f()
    }
}

解析
考点：闭包延迟求值
for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。 结果：

0xc042046000 2
0xc042046000 2

如果想不一样可以改为：

func test() []func()  {    
varfuns []func()
    fori:=0;i<2;i++  {
        x:=i
        funs = append(funs,func() {           
println(&x,x)
        })
    }    returnfuns
}

29.编译执行下面代码会出现什么?

package main
functest(x int)(func(),func()) {    
returnfunc() {       
println(x)
    x+=10
    }, func() {       
      println(x)
    }
}
funcmain() {
    a,b:=test(100)
    a()
    b()
}

解析
考点：闭包引用相同变量*
结果：

100
110

30. 编译执行下面代码会出现什么?

package main
import (   "fmt"
    "reflect")
funcmain1() {    
deferfunc() {     
iferr:=recover();err!=nil{
          fmt.Println(err)
       }else {
          fmt.Println("fatal")
       }
    }()    
deferfunc() {       
panic("deferpanic")
    }()    
panic("panic")
}
funcmain() {    
deferfunc() {       
iferr:=recover();err!=nil{
           fmt.Println("++++")
           f:=err.(func()string)            
fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
        }else {
           fmt.Println("fatal")
        }
    }()    
deferfunc() {       
panic(func()string {           
return "defer panic"
        })
    }()    
panic("panic")
}

解析
考点：panic仅有最后一个可以被revover捕获
触发panic("panic")后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的panic("panic")



package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

line := inpit.Text()
counts[line] = counts[line] + 1

input := buffio.NewScanner(os.Stdin)

package main 

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",Handle)
	log.Fatal(http.ListenAndServer("localhost:8080",nil))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}


f, err := os.Open(name)
if err != nil {
	return err
}
f.Close()

in, err := os.Open(file)
// ...
out, err := os.Crate(outfile)

p := new(int)
fmt.Println(*p)

age := make(map[string]int)

ages := map[string]int {
	"alice": 31,
	"charlie": 34
}

ages["bob"] = ages["bob"] + 1
ages["bob"] += 1
ages["bob"] ++


for name, age := range ages {
	fmt.Printf("%s\t%d\n", name, age)
}

import "sort"

var names []string
for name := range ages {
	names = append(names,name)
}
sort.Strings(names)
for _, name := range names {
	fmt.Printf("%s\t%d\n",name, ages[name])
}

gopl.io/ch5/title1
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html",url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title"&&n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func title(url string) err {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return fmt.Errorf("%s has type %s, not text/html",url,ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	// ...print doc's title element
}
io/ioutil
package ioutil
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}

// 互斥锁
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
Reader/ Closer




















