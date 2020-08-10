Go语言编程开发过程中需要警惕或者避开的一些坑：

    - 切片append添加的是单个元素不是切片：
    - goroutine创建的顺序不代表是执行的顺序

    - select 机制：
        - select+case用于阻塞监听goroutine
        - select底下有多个可执行的case，则随机执行一个
        - select常配合for循环来监听channel有没有故事发生，其中break只是退出当前select

    - defer机制：
    -       defer先进后出的方式执行
    -       defer在for循环中可能导致性能问题
    -       defer一定要在判断err之后
    -       调用os.Exit后defer不会被执行
    -       panic遇到defer后先执行defer再传递panic

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

1、olang中defer、return、返回值之间执行顺序的坑