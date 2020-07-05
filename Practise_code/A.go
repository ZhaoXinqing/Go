type Locker interface {
	Lock()
	Unlock()
}
type Once struct {

}

func (o *Once) Do(f func())

config.once.Do(func() { config.init(filename)})

type Mutex

Dial函数和服务端建立连接：
conn, err := net.Dial("tcp", "gool.com:80")
if err != nil {
	// handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReaderString('\n')
// ...

Listen函数创建的服务端：
ln, err :=net.Listen("tcp", ":8080")
if err != nil {
	// hand error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
		continue
	}
	go handleConnection(conn)
}

os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取。
file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
// 若打开失败，open file.go: no such file or directory

文件的信息可以读取进一个[]byte切片。Read和Write方法从切片参数获取其内的字节数。
data := make([]byte,100)
count,err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])

import "path"
path实现了对斜杠分隔的路径的实用操作函数。

Logger会打印每条日志信息的日期、时间，默认输出到标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。
Panic系列函数会在写入日志信息后panic。

import "flag"
flag包实现了命令行参数的解析。
要求：使用flag.String(), Bool(), Int()等函数注册flag，下例声明了一个整数flag，解析结果保存在*int指针ip里：
var ip = flag.Int("flagname",1234,"help message for flagname")
如果你使用的是flag自身，它们是指针；如果你绑定到了某个变量，它们是值。
fmt.Println("ip has value", *ip)
fmt.Println("flagver has value", flagver)
将flag绑定到一个变量，使用Var系列函数：
var flagvar int
func init() {
	flag.InVar(&flagver, "flagname", 1234, "help message for flagname")
}

import "flag"

var ip = flag.Int("flagname", 1234, "help message for flagname")
var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}

errors包实现了创建错误值的函数。
package error_test
import (
	"fmt"
	"time"
)

//MyError is an error implementation that includes a time and message.
type MyError struct {
	when time.time
	what string
}
func (e MyError) Error() string {
	return fmt.Sprint("%v:%v",e.when,e.what)
}
func opps() error {
	return MyError{
		time.Data(1989, 3, 15, 22, 30, 0, 0, time.UTC)
		"the file system has gone away"
	}
}
func Example() {
	if err != opps(); err != nil {
		fmt.Println(err)
	}
	// Output:
}
func (h Hash) Size() int
func (h Hash) Availlable() Bool

import "container/list"
//list 包实现了双向链表。要遍历一个链表：
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}

//Creat a new list and put some numbers in it.
l := list.New()
e4 := l.PushBack(4)
e1 := l.PushFront(1)
l.InterBefore(3, e4)
l.InterAfter(2, e1)
// Iterate through list and print its contants.
for e := l.Front(); e != nil; e = e.Next() {
	fmt.Println(e.Value)
}
tem1, err := template.New("name").Parse(...)
// 省略错误检测
err = tmpl.Execute(out, data)


http包提供了HTTP客户端和服务端的实现。
resp, err := http.Get("http://example.com")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",url.Values{"key":{"value","id":{"123"}}})

resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
//...

要管理服务端的行为，可以创建一个自定义的Server：
s := &http.Server{
	Addr: ":8080",
	Handle: myHandler,
	ReadTimeout: 10*time.Second,
	writeTimeout: 10 * time.Second,
	Max.Fatal(s.listenAndServer())
}
log.Fatal(s.listenAndServer())

//syslog 包提供一个简单的系统日志服务的接口。 It can send messages 
//to the syslog daemon using UNIX domain sockets, UDP or TCP.
sysLog, err := syslog.Dial("tcp", "localhost:1234", 
	syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
if err != nil {
	log.Fatal(err)
}
fmt.Fprintf(syslog,"This is a daemon warning with demotag.")
sysLog.Emerg("And this is a daemon emergency with demotag.")

/* 
reflect包实现了运行时反射，允许程序操作任意类型的对象。
典型用法是用静态类型interface{}保存一个值，
通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。
调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。
Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
*/

type S struct {
	F string 'species:"gopher" color:"blue"'
}

s := S{}
st := reflect.TypeOf(s)
field := st.Field(0)
fmt.Println(field.Tag.Get("color"),field.Tag.Get("species"))

// 判断字符串s是否包含字符串chars中的任一字符。
func ContainsAny(s, chars string) bool

// 返回字符串s中有几个不重复的sep子串。
func Count(s, sep string) int

// 返回将所有字母都转为对应的小写版本的拷贝。
func ToLower(s string) string
fmt.Println(string.Tolower("Gopher"))

/* Match检查b中是否存在匹配pattern的子序列。MatchString类似Match，
但匹配对象是字符串。*/

func MatchString(pattern string, s string) (matched bool, err error)

matched, err := regexp.MatchString("foo", "seafood")

/*Go 没有类。然而，仍然可以在结构体类型上定义方法。
方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

/*Go 没有类。然而，仍然可以在结构体类型上定义方法。
方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f > 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
// &{15 20} 25

/* 接口类型是由一组方法定义的集合。
接口类型的值可以存放实现这些方法的任何值。
*/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f 
	a = &v 

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 5

// 隐式接口解藕了实现接口的包和定义接口的包：互不依赖。
package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

// 一个普遍存在的接口是 fmt 包中定义的 Stringer。
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)

package main

import "fmt"

type IPAddr [4]byte

func main() {
	addrs := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
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


package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// 修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值
package main

import "fmt"

func Sqrt(x float64) (float64, error) {
	return 0,nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// 创建一个 strings.Reader。 并且以每次 8 字节的速度读取它的输出
// Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 
// 在遇到数据流结尾时，返回 io.EOF 错误。
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n",b[:n])
		if err == io.EOF {
			break
		}
	}
}

// 实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

func main() {
	reader.Validate(MyReader{})
}

/* 包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r * Request)
}

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "Hello!")
	}
)

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000",h)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}

package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
)

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(32, 56))
}

// 函数可以返回任意数量的返回值。swap 函数返回了两个字符串。
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

var a, b string

func main() {
	a, b = swap("hello", "zhaoxinqing")
	fmt.Println(a,b)
}
// zhaoxinqing hello

/* 命名返回值
Go 的返回值可以被命名，并且像变量那样使用。
返回值的名称应当具有一定的意义，可以作为文档使用。
没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum *4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

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


func main() {
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