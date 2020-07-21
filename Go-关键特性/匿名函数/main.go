
1、GO语言的匿名函数就是闭包

基本概念：
　　闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。
闭包的价值：
　　闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。

Go语言中的闭包

　　Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。例：

// closure.go
package main

import (
    "fmt"
)

func main() {
    var j int = 5

    a := func() func() {
        var i int = 10
        return func() {
            fmt.Printf("i, j: %d, %d\n", i, j)
        }
    }()　　//末尾的括号表明匿名函数被调用，并将返回的函数指针赋给变量a

    a()

    j *= 2

    a()
}
输出：

i, j: 10, 5
i, j: 10, 10

2、Go语言支持匿名函数，即函数可以像普通变量一样被传递或使用。

闭包是“函数”和“引用环境”组成的整体。

func anonymous(n int) func() {
    return func() {
        n++ //对外部变量加1
        fmt.Println(n)
    }
}

func anonymous2(n int) func() {
    sum := n
    a := func() { //把匿名函数作为值赋给变量a(Go不允许函数嵌套，但你可以利用匿名函数实现函数嵌套)
        fmt.Println(sum + 1) //调用本函数外的变量，这里没有()匿名函数不会马上执行
    }
    return a
}

　　fmt.Println("---------anonymous func--------")
    anony := anonymous(10)
    anony()
    anony1 := anonymous(20)
    anony1()
    /**
    *再次调用anony()函数，结果是12，由此得出以下两点
    * 1、内函数对外函数的变量的修改，是对变量的引用
    * 2、变量被引用后，它所在的函数结束，该变量不会马上销毁
    **/
    anony()
    anony1()
    fmt.Println("---------anonymous2----------")
    an := anonymous2(10)
    an()
    an2 := anonymous2(20)
    an2()
    an()
    an2()
输出：

---------anonymous func--------
11
21
12
22
---------anonymous2----------
11
21
11
21

闭包函数出现的条件：
1.被嵌套的函数引用到非本函数的外部变量，而且这外部变量不是“全局变量”;
2.嵌套的函数被独立了出来(被父函数返回或赋值 变成了独立的个体)，  而被引用的变量所在的父函数已结束。
对象是附有行为的数据，而闭包是附有数据的行为。

3、匿名函数执行

　　匿名函数最后括号中加参数，匿名函数立即执行，没有括号或括号中为空，则需要传参。例：

fmt.Println("-----------匿名函数的执行----------")
    m, n := func(i, j int) (m, n int) { // x y 为函数返回值
        return j, i
    }(1, 9) // 直接创建匿名函数并执行

    fmt.Println(m, n)

    f := func(i, j int) (result int) { // f 为函数地址
        result = i + j
        return result
    }

    fmt.Println(f(1, 2))
输出：

-----------匿名函数的执行----------
9 1
3
4、错误： expression in go/defer must be function call

　　go关键字后跟匿名函数报错： expression in go/defer must be function call，在匿名函数后加括号则可。例：

func smtp() error {
    return nil
}
　　go func() {
        err := smtp()
        if err != nil {
            fmt.Printf(err.Error())
        }
    }()
　　若去掉上面代码中的空括号会报错。

　　加()封装成表达式。

　　
————————————————
版权声明：本文为CSDN博主「芸复山人」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ycy258325/java/article/details/54632915