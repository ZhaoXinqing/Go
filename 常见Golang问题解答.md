第一章、常见问题
1、Golang并发控制方法
在golang并发行为控制中，有三种常见的方式，分别是WaitGroup、channel和Context。
列举的三种Golang中并发行为控制模式。模式之间没有好坏之分，只在于不同的场景用恰当的方案。实际项目中，往往多种方式混合使用。
WaitGroup：位于sync包下，多个goroutine的任务处理存在依赖或拼接关系；
Channel+select：可以主动取消groutine；多goroutine中数据传递；channel可以代替WaitGroup的工作，但会增加代码逻辑复杂性；多channel可以满足Context的功能，同样，也会让代码逻辑变得复杂。
Context：多层级groutine之间的信号传播（包括元数据传播，取消信号传播、超时控制等）。

1、WaitGroup（当某任务需要多任务协同工作时，采用此方式）
func main() {
    var wg sync.WaitGroup

wg.Add(2) //添加需要完成的工作量2

    go func() {
        wg.Done() // 完成工作量1
        fmt.Println("goroutine 1 完成工作！")
}()

    go func() {
        wg.Done() // 完成工作量1
        fmt.Println("goroutine 2 完成工作！")
}()

    wg.Wait() // 等待工作量2均完成
    fmt.Println("所有的goroutine均已完成工作")
}

2、Channel+select（及时通知关闭不需要的goroutine，防止空转，造成泄露，）
func main() {
    exit := make(chan bool)

    go func() {
        for {
            select {
            case <-exit:
                fmt.Println("退出监控")
                return
            default:
                fmt.Println("监控中")
                time.Sleep(2 * time.Second)
            }
        }
}()

    time.Sleep(5 * time.Second)
    fmt.Println("通知监控退出")
exit <- true

    //防止mian gorotine 过早退出
    time.Sleep(5 * time.Second)
}

3、Context（解决启动后不可控给的问题）
func A(ctx context.Context, name string) {
    go B(ctx, name) //A 调用了B
    for {
        select {
        case <-ctx.Done():
            fmt.Println(name, "A退出")
            return
        default:
            fmt.Println(name, "A do someting")
            time.Sleep(2 * time.Second)
        }
    }
}

func B(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println(name, "B退出")
            return
        default:
            fmt.Println(name, "B do something")
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
ctx, cancel := context.WithCancel(context.Background())

go A(ctx, "【请求1】") // 模拟clint来了一个连接请求

    time.Sleep(3 * time.Second)
    fmt.Println("clent断开连接，通知对应处理client请求的A，B退出")
cancel() // 假设满足某条件client断开了连接，那么就传播取消信号，ctx.Done()中得到取消信号

    time.Sleep(3 * time.Second)
}


2、并发安全的的sync.map
如何选择Map
从性能测试结果可以看出，sync.Map并不是为了代替锁+map的组合。它的设计，是为了在某些并发场景下，相对前者能有较小的性能损耗。
两种情况下应该选择sync.Map
1．key值一次写入，多次读取（即写少读多场景）。
2．多个goroutine的读取、写入和覆盖在不相交的key集。

3、静态类型和动态类型
静态类型（static type），变量声明的时候的类型。
var age int // int 是静态类型
var name string // string 也是静态类型
动态类型（concrate type，也叫具体类型）是程序运行时系统才能看见的类型。（空接口可以承接任何类型的数据）
var i interface（）
i = 18
i = "Go编程时光"

接口类型，每个接口变量，实际上都是一pair对组合而成，pair对中记录这实际变量的值和类型。
var age int = 25
知道了接口的组成后，我们在定义一个变量时，除了使用常规的方法，也可以使用如下的方式。
func main() {
    age := (int)(25)
    //或者使用 age := (interface{})(25)

    fmt.Printf("type:%T,data:%v", age, age)
}
根据接口是否包含方法，可以将接口分为iface和eface
type Phone interface { // iface,带有一组方法的接口
    call()
}

var i interface{} //eface,不带有方法的接口


4、反射的必要性
由于动态类型的存在，在一个函数中接受的参数的类型有可能无法预先知晓，此时我们就要对参数进行反射，然后根据不同的类型做不同的处理。

5、创建和声明变量的五种方法
第一种：一行声明一个变量
var <name> <type>
声明过程中初始化
var name string = "Go编程时光" // Go中双引号表示字符串，单引号便是rune类型字符，和python时不一样的额
var name = "Go编程时光"  //上面声明的简化（小数不能简化，因为默认的float64精度很大，占内存）

第二种：多个变量一起声明
var (
    name string
    age int
    gender string
)

第三种：声明和初始化一个变量（用于函数内部的短声明法）
name := "Go编程时光"
//等价于
var name string = "Go编程时光"
//等价于
var name = "Go编程时光"

第四种：声明和初始化多个变量
name,age := "Go编程时光"，28
//常可用于变量的交换
var a int = 100
var b int = 200
b, a = a, b

第五种：new函数创建指针变量
两种变量：存放数据本身的是普通变量，存放数据地址的是指针变量。
var age int = 28
var ptr = &age  //&后跟变量名，表示取出该变量的内存地址


new函数
（Go里的一个内建函数）(是一个语法糖？)
使用表达式new(type)将创建一个Type类型的匿名变量，初始化为type类型的零值，然后返回变量地址，返回的指针类型是*Type
func main() {
    ptr := new(int)
    fmt.Println("Ptr address: ",ptr)  //ptr address: 0xc000010098
    fmt.Println("ptr vakue:", *ptr)  // (*后面接指针变量，表示从内存地址中取出值) ptr value: 0
}
//使用传统的方式
func newInt() *int {
    var dummy int
    return &dummpy
}
// 使用new
func newInt() *int {
    return new(int)
}


匿名变量（优点有三）
不分配内存，不占用内存空间
不需要你为命名无用的变量名纠结
多次声明不会用任何问题（其它变量/常量只能声明一次，声明多次，编译就会报错）
通常我们用匿名接受必须接受，但是又不会用到的值。
func GetData() (int, int) {
    return 100,200
}
func main() {
    a, _ := GetData()
    _, b := GetData()
    fmt.Println(a, b)
}

6、Golang中还没有的泛型
使用泛型能够减少重复的代码和逻辑，为工程师提供更强的表达能力从而提升效率。
为什么Go到现在为之也不支持泛型？
泛型困境使我们必须在开发效率、编译速度和运行速度三者中选择两个；
目前社区中的Go语言方案都是有缺陷的，而Go团队认为泛型的支持不够紧急；

7、

9、map如何顺序读取
map不能顺序读取，是因为他是无序的，想要有序读取，首先的解决的问题就是，把ｋｅｙ变为有序，所以可以把key放入切片，对切片进行排序，遍历切片，通过key取值。

