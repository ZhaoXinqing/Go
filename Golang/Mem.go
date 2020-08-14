// Go语言安全研究-内存安全

// 重释放：
	// 指针重释放：Go使用垃圾收集器自动管理内存，无需显示释放内存
	// 资源冲释放：重复释放go对象资源会触发运行时错误
	
// 空指针：
	// 解引用一个空指针能导致Go程序崩溃，确保对象在使用前被有效构建并初始化好

// 内存分配大小控制：
	// 对输入和计算后new分配的数组大小和make slice、map、chan对象的大小进行合法性校
	// Mheap、

// unsafe包引发的内存风险
	// 指针在Go中是“安全”的，指针在正常的代码编写中无法指向任意的内存区域
	// 使用unsafe包的Unsafe.Pointer和Uintptr方法绕过指针使用的限制，则会出问题；
		// 访问越界：
		// 缓冲区溢出：
		// 类型转换：unsafe.Pointer可以把指针转换成任意类型，若转换后的类型过大，非法区域将被越界访问

// 内存泄漏：
	// 使用切片时：底层数组较大，而切片引用范围较小，因被引数组无法被GC回收，造成内存占用易导致内存泄漏；
		// 解决办法：如果切片的原始数据较大且切片的范围较小，则对返回的切片从原数据中复制一份出来，
		// 脱离原数据，使原数据不再被其他任何地方引用而被能被GC感知并收集
	// Goroutine产生内存泄漏：（设置安全退出机制，例如超时退出）
		// goroutine本身的栈所占用的空间造成内存泄露。
		// goroutine中的变量所占用的堆内存导致堆内存泄露.
	// 不正确使用runtime.SetFinalizer（指针构成的 "循环引用" 加上 runtime.SetFinalizer 会导致内存泄露）
		// 垃圾回收器能正确处理 "指针循环引用"，但无法确定 Finalizer 依赖次序，
		// 也就无法调用Finalizer 函数，这会导致目标对象无法变成不可达状态，其所占用内存无法被回收。
		// 禁止SetFinalize和指针循环引用同时使用

// 示例代码
package main
 
import(
	"fmt"
	"time"
	"runtime"
)
 
func main(){
	fmt.Println("内存安全")
	
	
//1、重释放
//重释放指针资源  无相关语法 不涉及 有GC自动回收机制
//不涉及 没有释放指针内存的语法  Go使用垃圾收集器自动
//管理内存，减少很多使用指针上的操心，首先我们再也不需
//要显示释放内存，悬挂指针（dangling pointer，指向已释
//放的内存）以及多次释放同一个指针指向内存的问题就不会
//发生，甚至不用担心内存是在栈上分配的还是堆上分配的，取
//一个变量的地址是安全的，不会产生悬挂指针。如果在程序中
//取了一个变量的地址，编译器会自动在栈上分配这个变量
 
//试验中调用的函数内部创建了很多指针对象，结束后并没有通过
//类似C++的free进行释放,退出函数后内存被回收。所以无需也没有释放语句
//从而导致指针重释放的问题
 

traceMemStats()
mutiFreeTest()
//runtime.GC()//强制启动GC垃圾回收
traceMemStats()
go func() {
	for {
		traceMemStats()
		time.Sleep(10 * time.Second)
	}
}()
	

 
//重释放资源 可能defer调用close重释放、条件判断没有理清多处close释放
//重复释放一般存在于错误处理流程判断中，如果恶意攻击者构造出错误条件使
//程序重复释放channel，则会触发运行时错误，从而造成DoS攻击。 重复close
//文件对象，在特殊的os的file实现中也可能会导致问题。对sql.DB对象和网络
//连接的close同样如此。 因此禁止重复释放资源	 
//panic: close of closed channel  exit status 2 

var Dividend rune 
var Divisor rune
var channel chan int = make(chan int)
fmt.Scanln(&Dividend,&Divisor)
result:=mutiClose(Dividend,Divisor,channel)
fmt.Printf("result is :%d",result)

	
//2、解引用空指针-----程序崩溃:runtime error: invalid memory address or nil pointer dereference

var varPtr *int;
//使用前进行判空 否则解引用空指针 程序崩溃
//if varPtr==nil {
//	varPtr = new(int) 
//}
*varPtr = 10;
fmt.Println(*varPtr)

 
 
 
//3、内存申请大小异常
//确保对输入和计算后make slice、map、chan对象的大小进行合法性校验
//申请过量内存  若外部数据无校验 panic: runtime error: makeslice: len out of range

var size uint64;
fmt.Scanln(&size)
arr:=make([]int,size)
fmt.Println("分配成功")
fmt.Println(arr[10])

 
//4、unsafa包操作内存 不严格的预防措施容易造成例如越界访问、缓冲区溢出等致命漏洞
//整数安全中的例子二可以反映
 
 
 
//5、内存泄漏
//切片使用引起的原数据的内存泄漏
//对于字符串来说，引用切片需要复制一份出来，解除对原字符串的引用
//  例如s0 = string([]byte(s1[:50]))   
//s0 = append([]int(nil), s1[len(s1)-30:]...)
 

traceMemStats()
//sliceLeak001()
sliceArr:=sliceLeak002()
//强制启动GC垃圾回收 查看内存是否泄露
runtime.GC()
traceMemStats()；
fmt.Println(*sliceArr[0])

 
 
//goroutine引起的内存泄漏
//确保每个协程都能退出
//协程（Goroutine）是Go语言并行设计的核心，启动一个协程就会做一个入栈操作，
//在系统不退出的情况下，协程也没有设置退出条件，则相当于协程失去了控制，
//它占用的资源无法回收，可能会导致内存泄露
//下面代码启动了两个协程，每个协程都是循环向屏幕上打印信息，在main()不退出的情况，且协程
//也没有设置退出条件，则导致协程所占用的资源以及启动协程的栈信息无法得到释放
outCh := make(chan int)
// 死代码 永不读取 对于无缓冲的chan变量 在向chan变量写数据后会阻塞
go func() {
	if false {
		<-outCh
	}
	select {}
}()
traceMemStats()
// 每s起100个goroutine，goroutine会阻塞，不释放内存
tick := time.Tick(time.Second / 100)
i := 0
for range tick {
	i++
	//fmt.Println(i)
	alloc1(outCh)
	if i==100{
		fmt.Println("done")
		break
	}
}

runtime.GC()
traceMemStats()

 
//6、禁止SetFinalize和指针循环引用同时使用以防止内存泄露

traceMemStats()
    for {
        test()     
		runtime.GC()//强制启动GC垃圾回收
		traceMemStats()
		time.Sleep(time.Millisecond)
    }
	
 
	
time.Sleep(1 * time.Hour) // 保持程序不退出	
}
 
func mutiFreeTest(){
	var bigArr [13107200] (*int)  //200M容量
	for i:=0;i<13107200;i++{
		var temp *int = new(int);
		bigArr[i] = temp
	}
	traceMemStats()
	
}
 
 
func mutiClose(Dividend rune,Divisor rune,channel chan int) rune {
	defer close(channel)
	
	//some operate 
	
	if(Divisor==0){
		close(channel)
		return 0
	}
	return Dividend/Divisor
}
 
 
//不返回切片 这样使得GC对数组收集
func sliceLeak001() {
	var bigArr [13107200] (*int)  
	for i:=0;i<13107200;i++{
		var temp *int = new(int);
		*temp = i
		bigArr[i] = temp
	}
	traceMemStats()
}
 
//返回切片一部分，使得其他部分不能够在此被使用且也没办法回收，产生内存泄露
func sliceLeak002() [](*int) {
	var bigArr [13107200] (*int)  
	for i:=0;i<13107200;i++{
		var temp *int = new(int);
		*temp = i
		bigArr[i] = temp
	} 
	ret:=bigArr[len(bigArr)-30:]
	traceMemStats()
	return ret
}
 
func alloc1(outCh chan<- int) {
    go func() {
        defer fmt.Println("alloc-fm exit")
        // 分配内存 协程中的堆内存也会同时泄漏
        buf := make([]byte, 1024*1024*10)
        _ = len(buf)
        //fmt.Println("alloc done")
 
        outCh <- 0
		fmt.Println("没有阻塞")
    }()
}
 
type Data struct {
    d   [1024 * 100]byte
    o   *Data
}
func test() {
    var a, b Data
    a.o = &b
    b.o = &a
    //runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
    //runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}
 
func memstate(){
	traceMemStats()
	sliceArr:= make([]int,8)
	for i:=1;i<=64*1024*1024;i++{
		sliceArr = append(sliceArr,i)
	}
	traceMemStats()
}
 
 
func traceMemStats() {
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
	//Alloc应用层申请的堆内存 HeapSys从操作系统申请的内存 HeapReleased返还给操作系统的内存
    fmt.Printf("Alloc:%d(bytes) HeapSys:%d(bytes) HeapReleased:%d(bytes)\n", 
				ms.Alloc, ms.HeapSys, ms.HeapReleased)
}