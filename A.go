import "container/list"
//list 包实现了双向链表。要遍历一个链表：
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}

//与 fmt.Stringer 类似，`error` 类型是一个内建接口：
type error interface {
	Error() string
}
//通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 `nil`， 
//来进行错误处理。error 为 nil 时表示成功；非 nil 的 error 表示错误
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't conver number: %v\n", err)
}
fmt.Println("Converted integer:", i)


// Go 的基本类型有Basic types
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint64 uintptr

byte // uint8 的别名
rune // int32 的别别名
	// 代表一个Unicode码

float32 float64

complex64 complex128


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