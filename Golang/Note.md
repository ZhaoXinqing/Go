go module 
go get
declarationimport package
varuables
Variables
Comments
set 
struct
interface 方法的集合
reflect.TypeOf(),reflect.ValueOf()
sync.map
map自动库扩容
ok,delete

Pointer
& *
cha

子切片也发生变化，重新分配内存【2：5]从2到5
但不包括5
copy(number5,numbers[:2])

name := make([]byte,length,cpacity)

值传递和引用传递，个数，字符串的长度
做函数作为实参；
实现闭包。匿名函数
方法，包含了一个接收者的函数。
递归函数
init函数
循环
for continue break switch,fallthrough

引用类型：array，slice，map, channel, interface;
return 

(1)使用new()、make()来创建变量，编译器会根据变量的大小及其"escapeanalysis"的结果来决定变量的存储位置（堆或栈），故能准确返回本地变量的地址，在gobuild或gorun时，加入-m参数，能准确分析程序的变量分配位置；

4、Golang的字符串本质是字符数组([]byte)：
(1)string类型的值是常量，不可更改。若修改字符串中的字符，将string转为[]byte修改后，再转为string即可。
(2)string的值不必是UTF8文本，可以包含任意的值。只有字符串是文字字面值时才是UTF8文本，字串可以通过转义来包含其他数据。判断字符串是否是UTF8文本，可使用"unicode/utf8"包中的ValidString()函数。

// 2、错误处理和抛出异常：defer，panic，recover：
// (1)多个defer的执行顺序为“后进先出”。defer、return、返回值三者的执行逻辑应该是：
	// ①return最先执行，return负责将结果写入返回值中；
	// ②接着defer开始执行一些收尾工作；
	// ③最后函数携带当前返回值退出。
// (2)defer延迟执行的函数，它的参数会在声明时候就会求出具体值，而不是在执行时才求值。
// golang的所有错误都是通过 error 来处理的：

(1)goroutine中操作是顺序执行的，多goroutine想顺序执行，可以使用channel或sync包中的锁机制；

recover() panic
nil
slice

索引，值
slice
case breaskfallthrough
1、程序退出时还有 goroutine 在执行，程序默认不等所有 goroutine 都执行完才退出。解决办法：使用 "WaitGroup" 变量，它会让主程序等待所有 goroutine 执行完毕再退出。

reflect deepequal()
13、更新 map 中 struct 元素的字段值，有 2 个方法：使用局部变量；使用指向元素的 map 指针；
12、 map 中的元素是不可寻址的，slice 的元素可寻址；

7、从 slice 中重新切出新 slice 时，新 slice 会引用原 slice 的底层数组。如果跳了这个坑，程序可能会分配大量的临时 slice 来指向原底层数组的部分数据，将导致难以预料的内存使用。可以通过拷贝临时 slice 的数据，而不是重新切片来解决。
8、for 语句中的迭代变量在每次迭代中都会重用，即 for 中创建的闭包函数接收到的参数始终是同一个变量，在 goroutine 开始执行时都会得到同一个迭代值，最简单的解决方法，在 for 内部使用局部变量保存迭代值，再传参，另一个解决方法：直接将当前的迭代值以参数形式传递给匿名函数。

8、在一个长时间执行的函数里，内部 for 循环中使用 defer 来清理每次迭代产生的资源调用，就会出现问题，解决办法：defer 延迟执行的函数写入匿名函数中。

7、从一个现有的非 interface 类型创建新类型时，并不会继承原有的方法，如果需要使用原类型的方法，可将原类型以匿名字段的形式嵌到你定义的新 struct 中；
8、 interface 看起来像指针类型，但它不是。interface 类型的变量只有在类型和值均为 nil 时才为 nil；
9、生产者消费者模型
(1)通过chan在生产者和消费者之间传递数据(ch)和同步状态(done)；
(2)chan作为参数传递时是引用传递，不需要使用指针；
(3)chan是协程安全的，多个goroutine之间不需要锁；
(4)chan的close事件可以被recv获取，close事件一定在正常数据读完之后，机制类似于read到EOF；

1、使用传址方式为 WaitGroup 变量传参。
2、若函数 receiver 传参是传值方式，则无法修改参数的原有值，除非 receiver 参数是 map 或 slice 类型的变量，并且是以指针方式更新 map 中的字段、slice 中的元素的，才会更新原有值。
3、Go 提供了一些库函数来比较那些无法使用 == 比较的变量，比如使用 "reflect" 包的 DeepEqual() ；
4、可以使用相等运算符 == 来比较结构体变量，前提是两个结构体的成员都是可比较的类型；
5、在 encode/decode JSON 数据时，Go 默认会将数值当做 float64 处理；
6、在 range 迭代中，得到的值其实是元素的一份值拷贝，更新拷贝并不会更改原来的元素，即是拷贝的地址并不是原有元素的地址，如果要修改原有元素的值，应该使用索引直接访问。


    - 切片append添加的是单个元素不是切片：
    - goroutine创建的顺序不代表是执行的顺序

    - select 机制：
        - select+case用于阻塞监听goroutine
        - select底下有多个可执行的case，则随机执行一个
        - select常配合for循环来监听channel有没有事件发生，其中break只是退出当前select

    - defer机制：

    - 将格式好的字符串输出
    -       Printf() 是把格式字符串输出到标准输出
    -       Sprintf() 是把格式字符串输出到指定字符串中，参数多一个char*，为目标字符串地址
    -       Fprintf() 是把格式字符串输出到制定文件设备中，参数多一个FILE*

    - nil map/slice 通过var声明没有字面量的map和slice后不会创建底层数组，无法直接添加元素 

    - array和struct是值类型的，而slice、map、chan、interface是引用类型，一般使用slice而不是array，struct也尽量使用指针，能够避免资源消耗和无法修改数据问题，特别注意对struct的for range行为

    - interface{}类型的变量无法调用属性，只能调用方法

    - struct属性如果是小写开头，则其序列化会丢失属性对应的值，同时也无法进行Json解析
    
    - 重写String()方法可能因为fmt选项原因 如%v 导致循环调用String()造成栈溢出

    - go语言触发异常的场景：空指针解析、下标越界、除数为0、调用panic函数

    - go语言中没有隐藏的this指针，因为方法施加的对象显式传递，没有被隐藏起来

    - make初始化是有默认值的 如s:=make([]int,3) s = append(s,1,2,3) 得到的是[0,0,0,1,2,3]

    - 在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名

    - struct中有map、slice属性时不能直接==比较

    - 常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，不能取地址

    - goto不能跳转到其他函数或者内层代码

    - type Myint1 int 和 type Myint2 int 不一样，第一个是创建一个新类型，底层是int，但不能直接=int，第二个是类型别名

    - 更新字符串字符，最好用rune，防止不用文字占用字节数不同导致无法预计的结果
    -       x := "text"
    -       xRunes := []rune(x)
    -       xRunes[0] = '我'
    -       x = string(xRunes)        x = "我ext"

    - go的内建函数len()返回的是字符串的byte数

    - 从一个现有的非interface类型创建新类型时，并不会继承原有的方法

    - map类型的key必须是可以使用==比较的类型，不能是slice、map、function。

    - 使用读写锁RWMutex时，如果有重入可能会导致死锁。

    - 函数传递切片，在函数中使用append后一定要把切片做一次返回，否则原切片不会改变。

    - 闭包中的外部变量不是锁死的，会随着外部变量改变。（好像指针一样）

    - 将for-select封装到函数中。

    - 初始化结构体时使用带标签的语法。

    - 将结构体初始化拆分为多行。

    - 利用 iota 来使用自定义的整数枚举类型，务必要为其添加 String() 方法。

    - 重度使用map操作数据，最好设置getter、setter等方法。

    - Go语言速度快主要归功于对依赖的管理。

    - 在string和[]byte之间转换，会给GC造成很大压力。底层数据结构会复制。
    
    - 使用+来进行string连接会生成新对象，降低GC效率。
