// Golang中的多态通过 interface（接口） 来实现的，
// 即定义一个接口类型，里面声明一些要实现的功能；
// 定义接口的所有方法的任何类型都表示隐式实现该接口。
// 类型接口的变量可以保存实现该接口的任何值。接口的这个属性用于实现GO的多态性。
type Person interface {SayHello() }

type Girl struct { Sex string }
type Boy struct { Sex string }

func (this *Girl) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}
func (this *Boy) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}

func main() {
	g := &Girl{"girl"}
	b := &Boy{"boy"}
	
	p := map[int]Person{}
	p[0] = g
	p[1] = b
 
	for _, v := range p {
		v.SayHello()
	} 
}


// 整形切片初始化
// s := make([]int) 错误
s := make([]int, 0)
s := make([]int, 5, 10)
s := []int{1, 2, 3, 4, 5}

// 正确channel语法
var ch chan int
ch := make(chan int)
// 声明一个只用于读取int数据的单项channel变量ch
var ch <-chan int

// slice和map操作
var s []int
s = append(s,1)

var s []int
s = make([]int,0)
s = append(s,1)

var m map[string]int
m = make(map[string]int)
m["one"] = 1

// 无缓冲的channel是同步的，而有缓冲的channel是异步的


// pointer

// 指针是C族语言（C/C++）极高性能的根本，在操作大块数据和做内存地址偏移时，方便快捷。
// C 族语言（C/C++）中指针被诟病的原因是 指针的运算 和 内存释放。（go语言去掉了指针运算，并能自动回收）
// go中的指针分为两个重要方面：
	// 类型指针。类型指针拥有指针的高效访问：传递数据使用指针，而不是采用内存拷贝；
			// 使用类型指针修改（其指针地址）指向的普通变量值。类型指针不能运算和偏移。
	// 切片。由指向起始元素的原始指针、元素数量和容量构成。
			// 切片比原始指针更安全高效，当切片越界时，go运行时会报panic，并打出堆栈，原始指针只能崩溃

// 1. 指针（变量）：取指针地址（&）和指针类型（*T）
	// & + 变量（取地址），指针值带有"0x"的十六进制前缀
	// * + 指针（取指针的值）
// 2. 使用指针修改值（常用）
	// t := *a，取指针a指向的变量值，赋给变量t 
	// *b = t，把普通变量t的值，赋给指针b指向的变量
// 3. 使用new( )函数创建指针（常用）
	// 创建过程会分配内存，指针指向的值为默认值


r := bufio.NewReader(f)
for {
	line, err := r.ReadString("\n")
	if err == io.EOF{
		break
	} else if err != nil {
		fmt.Println("error reading file %s\n",err)
	}
}

type Reader interface {
	ReadFrom(r Reader) (n int64,err error)
}

type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

var c1 chan string = make(chan string)
func() {
	time.Sleep(time.Second * 2)
	c1 <- "result 1"
}()
fmt.Println("c1 is",<-c1)
// 以上代码会死锁，因为push和pull永远不可能发生，
// 这是阻塞channel的不当用法

// 解决办法————1
var c1 chan string = make(chan string)
go func() {
	time.Sleep(time.Second * 2)
	c1 <- "result 1"
}()
fmt.Println("c1 is",<-c1)
// 通过在另一个协程中run push 代码，使得channel的生产和消费可以同时对接

// 解决办法————2
var c1 chan string = make(chan string，1)
func() {
	time.Sleep(time.Second * 2)
	c1 <- "result 1"
}()
fmt.Println("c1 is",<-c1)
// 给channel加一个buffer，只要buffer没用尽，大家就不会阻塞
