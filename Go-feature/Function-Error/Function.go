// 函数是最基本的代码块，用于执行一个任务。
// 可以通过函数来划分不同的功能
// 函数声明告诉了编译器函数的名称，返回类型，和参数；
// Go 语言标准库提供了多种可动用的内置的函数。例如，
// 		- len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，
// 		- 如果传入的是数组，则返回数组中包含的元素个数。

// 函数定义：
func function_name([parameter list]) [return_types] {
	函数体
}
// 函数定义解析：
// 		- func：函数由 func 开始声明
// 		- function_name：函数名称，函数名和参数列表一起构成了函数签名。
// 		- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
// 		- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
// 		- 函数体：函数定义的代码集合。
// 实例：max（）函数，函数返回两个数的最大值 
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int
	if (num1 > num2) {
		result = num1
	} else {
	   result = num2
	}
	return result
}

//  函数调用
func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int
	/* 调用函数并返回最大值 */
	ret = max(a, b)
	fmt.Printf( "最大值是 : %d\n", ret )
}

// 函数返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

// 函数参数
// 函数如果使用参数，该变量可称为函数的形参。
// 形参就是定义在函数体内的局部变量。
// 调用函数，可以通过两种方式来传递参数；
// 		- 值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
// 		- 引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
// 		- 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
// Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，
// 或者让方法能够改变接受的数据。

// 函数的用法：
// 		1、函数作为另外一个函数的实参，函数定义后可作为另外一个函数的实参数传入
// 		2、闭包，闭包是匿名函数，可在动态编程中使用
// 		3、方法，方法就是一个包含了接受者的函数；

// （1）函数作为实参
// 在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()
func main(){
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
	   return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}
// （2）函数实现闭包
// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
// 匿名函数的额优势在于可以直接使用函数内变量而不用声明；
// 实例中创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSequence() func() int {
	i:=0
	return func() int {
	   i+=1
	  return i  
	}
}
// ……待续
// Go 支持通过 闭包来使用 匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
// 由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包
// https://blog.csdn.net/ycy258325/article/details/54632915

// （3）方法
// go语言中同时有函数和方法，一个方法就是一个包含了接收者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针
// 所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
	/* 函数体*/
}
// 实例：定义一个结构体类型和该类型的一个方法
type Circle  struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
// 该method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius  // c.radius 即为 Circle 类型对象中的属性
}

// 递归函数
// 运行过程中调用自己，语法格式：
func recursion() {
	recursion() /* 函数调用自身 */
}
func main() {
	recursion()
}
// Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
// 递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
// 实现阶乘：
func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}
func main() {  
    var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
// 实现斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
	 return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
} 
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
//  fibonacci 实现，用到多返回值特性，降低复杂度：


// 函数补充：
// 定义一个包内全局字符串变量
var str string
var str = " "

// 通过指针变量p访问其成员变量 name
p.name
(*p).name

// 接口
// A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
// B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
// C. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口

// 字符串连接
str := "abc" + "123"
fmt.Sprintf("abc%d",123)

// 协程
// A. 协程和线程都可以实现程序的并发执行
// B. 通过channel来进行协程间的通信

// init函数
// A. 一个包中，可以包含多个init函数
// B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数

// 循环语句
// A. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
// B. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量 

func add(args ...int) int {
	sum := 0
	for _,arg := range args {
		sum += arg
	}
	return
}
// 调用：
add(1,2)
add(1,2,3)
add([]int{1,3,7}...)

// 类型转换
type Mylnt int
var i int = 1
var jMylnt = Mylnt(i)

// 局部变量初始化
var i int = 10
var i = 10
i := 10

// const常量定义
const Pi float64 = 3.14159265358979323846
const zero = 0.0

const (
	size int64 = 1024
	eof = -1
)

const u,v float32 = 0, 3
const a, b, c = 3,4,"foo"

// 布尔变量赋值
b = true
b = (1 == 2)

func main() {
	if (true){
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}

// switch
// 用break来明确退出一个case
// 在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case

// golang中没有隐藏的this指针
// A. 方法施加的对象显式传递，没有被隐藏起来
// B. golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达
// C. 方法施加的对象不需要非得是指针，也不用非得叫this

// golang的引用类型包括
// 数组切片、map、channel、interface

// golang中的指针运算
// 通过 “&” 取指针的地址
// 通过 “*” 取指针指向的数据

// main函数（可执行程序的执行起点）
// A. main函数不能带参数
// B. main函数不能定义返回值
// C. main函数所在的包必须为main包
// D. main函数中可以使用flag包来获取和解析命令行参数

// 赋值
var x interface{} = nil
var x error = nil

// 整形切片初始化
// s := make([]int) 错误
s := make([]int, 0)
s := make([]int, 5, 10)
s := []int{1, 2, 3, 4, 5}

// 切片中删除一个元素
func (s *Slice) Remove(value interface{}) error {
	for i,v := range *s {
		if isEqual(value, v) {
			*s = append((*s)[:i],(*s)[i + 1:]...)
			return nil
		}
	}
	return ERR_ELEM_NT_EXIST
}

// 变量的自增自减
i := 1
i++

i := 1
i--

// 函数声明
A. func f(a, b int) (value int, err error)
B. func f(a int, b int) (value int, err error)
// C. func f(a, b int) (value int, error) 错误
D. func f(a int, b int) (int, int, error)

// 若Add函数的调用代码为：
func main() {
	var a lnteger = 1
	var b lnteger = 2
	var i interface{} = &a
	sum := i.(*lnteger).Add(b)
	fmt.Println(sum)
}
// 则Add函数的调用代码：
type lnteger int
func (a lnteger) Add(b lnteger)lnteger{
	return a + b
}

type lnteger int
func (a *lnteger) Add(b lnteger) lnteger{
	return *a + *b
}

// 若Add函数的调用代码为：
func main() {
	var a lnteger = 1
	var b lnteger = 2
	var i interface{} = a
	sum := i.(lnteger).Add(b)
	fmt.Println(sum)
}
// 则Add函数的定义为：
type lnteger int
func (a lnteger)Add(b lnteger)lnteger{
	return a + b
}

// GetPodAction定义
type Fragment interface{
	Exec(transInfo *TransInfo)error
}
type GetPodAction struct {}
func (g GetPodAction)Exec(transInfo *TransInfo)error{
	...
	return
}
// 赋值正确的是
var fragment Fragment = new(GetPodAction)
var fragment Fragment = &GetPodAction{}
var fragment Fragment = GetPodAction{}

// GoMock
// A. GoMock可以对interface打桩
// B. GoMock打桩后的依赖注入可以通过GoStub完成

// 接口:
// A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
// B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
// C. 接口查询是否成功，要在运行期才能够确定

// 正确channel语法
var ch chan int
ch := make(chan int)
<- ch
// ch <-  错误

// 同步锁
// A. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
// B. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
// C. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占

// 指针数据类型可以转化为有效的JSON文本

// go vendor：
// A. 基本思路是将引用的外部包的源代码放在当前工程的vendor目录下面
// B. 编译go代码会优先从vendor目录先寻找依赖包
// C. 有了vendor目录后，打包当前的工程代码到其他机器的$GOPATH/src下都可以通过编译

// flag是bool型变量，表达式符合编码规范：
if flag == false
if ！flag

// value是整型变量，符合编码规范的是
if value == 0
if value != 0

// 函数返回值的错误设计
// A. 如果失败原因只有一个，则返回bool
// B. 如果失败原因超过一个，则返回error
// C. 如果没有失败原因，则不返回bool或error
// D. 如果重试几次可以避免失败，则不要立即返回bool或error

// 关于异常设计：
// A. 在程序开发阶段，坚持速错，让程序异常崩溃
// B. 在程序部署后，应恢复异常避免程序终止
// C. 对于不应该出现的分支，使用异常处理

// slice和map操作
var s []int
s = append(s,1)

var s []int
s = make([]int,0)
s = append(s,1)

var m map[string]int
m = make(map[string]int)
m["one"] = 1

// channel
// A. 给一个 nil channel 发送数据，造成永远阻塞
// B. 从一个 nil channel 接收数据，造成永远阻塞
// C. 给一个已经关闭的 channel 发送数据，引起 panic
// D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

// 无缓冲的channel是同步的，而有缓冲的channel是非同步的

// 关于异常的触发
// A. 空指针解析
// B. 下标越界
// C. 除数为0
// D. 调用panic函数

// cap函数的适用类型：
// array、slice、channel

// beego框架
// A. beego是一个golang实现的轻量级HTTP框架
// B. beego可以通过注释路由、正则路由等多种方式完成url路由注入
// C. 可以使用bee new工具生成空工程，然后使用bee run命令自动热编译

// goconvey
// A. goconvey是一个支持golang的单元测试框架
// B. goconvey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到web界面
// C. goconvey提供了丰富的断言简化测试用例的编写
// D. goconvey可与go test集成

// go vet
// A. go vet是golang自带工具go tool vet的封装
// B. go vet可以使用绝对路径、相对路径或相对GOPATH的路径指定待检测的包
// C. go vet可以检测出死代码

// map
// A. map反序列化时json.unmarshal的入参必须为map的地址
// B. 在函数调用中传递map，则子函数中对map元素的增加会导致父函数中map的修改
// C. 在函数调用中传递map，则子函数中对map元素的修改会导致父函数中map的修改
// D. bu能使用内置函数delete删除map的元素

1-100

// GoStub
// A. GoStub可以对全局变量打桩
// B. GoStub可以对函数打桩
// C. GoStub不可以对类的成员方法打桩
// D. GoStub可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为

// select机制
// A. select机制用来处理异步IO问题
// B. select机制最大的一条限制就是每个case语句里必须是一个IO操作
// C. golang在语言级别支持select关键字
// D（错误）. select关键字的用法与switch语句非常类似，后面要带判断条件

// 内存泄漏
// A. golang有自动垃圾回收，还是存在内存泄露
// B. golang中检测内存泄露主要依靠的是pprof包
// C. 内存泄露不可以在编译阶段发现
// D. 应定期使用浏览器来查看系统的实时内存信息，及时发现内存泄露问题

// 声明一个整形变量i
var i int

// 声明一个含有10个元素的整形数组a
var a [10]int

// 声明一个整形数组切片s
var s []int

// 声明一个整形指针变量p
var p *int

// 声明一个key为字符串型value为整形的map变量m
var m map[string]int

// 声明一个入参和返回值均为整形的函数变量f
var f func(a int) int

// 声明一个只用于读取int数据的单项channel变量ch
var ch <-chan int

// 源文件的命名为slice.go，则测试文件为：
// slice_test.go

// go test要求测试函数的前缀必须命名为： 
// Test

// 
for i :=0 ; i < 5; i++ {
	defer fmt.Printf("%d", i)
}
// 运行结果为：4 3 2 1 0 

// 
func main() {
	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Println(x)
}
// 运行结果：2 1

// 
func main() {
	strs := []string{"one","two","three"}

	for _,s := range strs {
		go func() {
			time.Sleep( 1 * time.Second)
			fmt.Printf("%s",s)
		}()
	}
	time.Sleep(3 * time.Second)
}
// 运行结果：three three three

func main() {
	x := []string{"a","b","c"}
	for v := range x {
		fmt.Print(v)
	}
}
// 运行结果：0 1 2

func main() {
	x := []string{"a","b","c"}
	for _,v := range x {
		fmt.Print(v)
	}
}
// 运行结果：a b c 

func main() {
	i := 1
	j := 2
	i,j = j,i
	fmt.Printf("%d%d\n",i,j)
}
// 运行结果：2 1

func incr(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	incr(&v)
	fmt.Println(v)
}
// 运行结果： 2

// 启动一个goroutine的关键字是：go

type Slice []int
func NewSlice() Slice {
	return make(Slice,0)
}
func (s *Slice) Add(elem int) *Slice{
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}
// 参考答案：1 3 2

// 数组是一个值类型

// 使用map不需要引入任何库

// 内置的delete不可以删除数组切片内的元素

// 指针不是基础类型

// interface{}是可以指向任意对象的Any类型

file, err := os.Open("test.go")
defer file.Close()
if err != nil {
	fmt.Println("open file failed:",err)
	return
}
// 文件操作可能触发异常

// golang支持垃圾自动回收

// golang可以复用C/C++模块，Cgo

// 从结构体实例编码到JSON数据格式的时候，
// 使用x作为名字，这可以看作是一种重命名的方式
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// 可以通过成员变量或函数首字母的大小写来决定其作用域

// golang支持goto语句

// 匿名函数可以直接赋值给一个变量或者直接执行

// 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
// 可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）

// 在函数的多返回值中，如果有error或bool类型，则一般放在最后一个

// 错误是业务过程的一部分，而异常不是

// 函数执行时，如果由于panic导致了异常，则延迟函数仍会执行

// 同级文件的包名不允许有多个

// golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承

// 使用for range 迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的。

// switch后面可以不跟表达式

// 结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
// 因此在decode时这些非导出变量的值为其类型的零值（）

// golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，
// 以NewXXX来命名

func deferDemo() error {
	err := creatResource1()
	if err != nil{
		return ERR_CREAT_RESOURCE1_FAILED
	}
	defer func() {
		if nil != nil{
			destroyResource1()
		}
	}()

	err = creatResource2()
	if err != nil {
		return ERR_CREAT_RESOURCE2_FAILED
	}
	defer func() {
		if err != nil {
			destroyResource2()
		}
	}()

	err = creatResource3
	if err != nil {
		return ERR_CREAT_RESOURCE3_FAILED
	}
	return nil
}

// channel 是支持只读只写单向传输的

func main() {
	defer_call()
}
func defer_call() {
	defer func(){fmt.Println("打印前")}()
	defer func(){fmt.Println("打印中")}()
	defer func(){fmt.Println("打印后")}()
	panic("触发异常")
}
// defer 是后进先出，panic需要等到defer结束后才会向上传递
// 打印后
// 打印中
// 打印前
// panic：触发异常

// 以下代码有什么问题
type student struct {
	Name string
	Age int
}
func pase_student() {
	m := make(map[string]*student)
	stus :=[]student{
		{Name:"zhou", Age:24},
		{Name:"li", Age:23},
		{Name:"wang", Age:22},
	} for _,stu := range stus {
		m[stu.Name] = &stu
	}
}
// for each 语句用于循环访问集合以获取所需信息，
// 但不应用于更改集合内容以避免产生不可预知的副作用

// 与Java的foreach一样，都是使用副本的方式。
// 所以m[stu.Name]=&stu实际上一致指向同一个指针， 
// 最终该指针的值为遍历的最后一个struct的值拷贝。 
// 就像想修改切片元素的属性：
for _,stu := range syus{
	stu.Age = stu.Age + 10
}
// 也是不可行的

// 正确写法
for i := 0; i<len(stus);i++ {
	m[stus[i].Name] = &stus[i]
}
for k,v := range m {
	println(k,"=v",v.Name)
}

// 代码会输出什么，并说明原因
func main() {
	runtime.GOMAXPROCESS(1)
	wg := sync.WaitGroup{}
	wg.Add(20) for i := 0;i<10;i++ {
		go func() {
			fmt.Println("A:",i)
			wg.Done()
		}()
	}
	for i := 0;i<10;i++ {
		go func(i int) {
			fmt.Println("B:",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
