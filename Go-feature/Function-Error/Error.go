// Golang没有try.. catch.. finally。所有的错误都是通过error来处理的
// 经典的代码如下:
err := xxxxxx
if err != nil {
    // ....
}
// 再高逼格的，就是把err在封装成各种类型。但基本都是这样一套处理方式。
// Golang也允许自己创建error信息，比如:
func Sqrt(value float64)(float64, error) {
    if (value < 0){
        return 0, error.New("Math:negative number passed to Sqrt")
    }
    return math.Sqrt(value)
}
// 用户可以通过errors.New()来创建自带业务逻辑的错误信息。



// Go语言通过内置的错误接口提供了非常简洁的错误处理机制
// err是一个接口类型，它的定义为：
type error interface {
	Error() string
}
// 我们可以在编码中通过error接口类型来生成错误信息。
// 函数通常在最后的返回值中返回错误信息。
// 使用errors.New 可返回一个错误信息，示例：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
}


// panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误，
// panic 用于主动抛出错误，recover 用来捕获 panic 抛出的错误。
// 引发panic有两种情况，一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。
// 发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，
// 然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数。
// panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。
// defer里面的panic能够被后续执行的defer捕获。
// recover用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，
// 但是defer只有在后面的函数体内直接被掉用才能捕获panic来终止异常，否则返回nil，异常继续向外传递。

// 示例
//以下捕获失败
defer recover()
defer fmt.Prinntln(recover)
defer func(){
    func(){
        recover() //无效，嵌套两层
    }()
}()
//以下捕获有效
defer func(){
    recover()
}()
func except(){
    recover()
}
func test(){
    defer except()
    panic("runtime error")
}

// 例子2
// 多个panic只会捕捉最后一个：
func main(){
    defer func(){
        if err := recover() ; err != nil {
            fmt.Println(err)
        }
    }()
    defer func(){
        panic("three")
    }()
    defer func(){
        panic("two")
    }()
    panic("one")
}
// 使用场景
// 一般情况下有两种情况用到：
//     - 程序遇到无法执行下去的错误时，抛出错误，主动结束运行。
//     - 在调试程序时，通过 panic 来打印堆栈，方便定位错误。

if result, errorMsg := Divide(100, 10); errorMsg == "" {
    fmt.Println("100/10 = ", result)
}

if _, errorMsg := Divide(100, 0); errorMsg != "" {
    fmt.Println("errorMsg is: ", errorMsg)
}
等价于:

result, errorMsg := Divide(100,10)
if errorMsg == ""{
    fmt.Println("100/10 = ", result)
}

result, errorMsg = Divide(100,0)
if errorMsg != ""{
    fmt.Println("errorMsg is: ", errorMsg)
}
GG
   GG

  401***8@qq.com

1年前 (2019-03-01)
   zhu波比

  568***900@qq.com

4
fmt.Println 打印结构体的时候，会把其中的 error 的返回的信息打印出来。

type User struct {
   username string
   password string
}

func (p *User) init(username string ,password string) (*User,string)  {
   if ""==username || ""==password {
      return p,p.Error()
   }
   p.username = username
   p.password = password
   return p,""}

func (p *User) Error() string {
      return "Usernam or password shouldn't be empty!"}
}

func main() {
   var user User
   user1, _ :=user.init("","");
   fmt.Println(user1)
}
结果：

Usernam or password shouldn't be empty!
zhu波比
   zhu波比

  568***900@qq.com

1年前 (2019-05-07)
   gibson1112

  185***72536@163.com

7
个人多次试验，总结几点 panic，defer 和 recover。

 1、panic 在没有用 recover 前以及在 recover 捕获那一级函数栈，panic 之后的代码均不会执行；一旦被 recover 捕获后，外层的函数栈代码恢复正常，所有代码均会得到执行；
 2、panic 后，不再执行后面的代码，立即按照逆序执行 defer，并逐级往外层函数栈扩散；defer 就类似 finally；
 3、利用 recover 捕获 panic 时，defer 需要再 panic 之前声明，否则由于 panic 之后的代码得不到执行，因此也无法 recover；
package main

import (
"fmt"
)

func main() {
  fmt.Println("外层开始")
  defer func() {
    fmt.Println("外层准备recover")
    if err := recover(); err != nil {
      fmt.Printf("%#v-%#v\n", "外层", err) // err已经在上一级的函数中捕获了，这里没有异常，只是例行先执行defer，然后执行后面的代码
    } else {
      fmt.Println("外层没做啥事")
    }
    fmt.Println("外层完成recover")
  }()
  fmt.Println("外层即将异常")
  f()
  fmt.Println("外层异常后")
  defer func() {
    fmt.Println("外层异常后defer")
  }()
}

func f() {
  fmt.Println("内层开始")
  defer func() {
    fmt.Println("内层recover前的defer")
  }()

  defer func() {
    fmt.Println("内层准备recover")
    if err := recover(); err != nil {
      fmt.Printf("%#v-%#v\n", "内层", err) // 这里err就是panic传入的内容
    }

    fmt.Println("内层完成recover")
  }()

  defer func() {
    fmt.Println("内层异常前recover后的defer")
  }()

  panic("异常信息")

  defer func() {
    fmt.Println("内层异常后的defer")
  }()

  fmt.Println("内层异常后语句") //recover捕获的一级或者完全不捕获这里开始下面代码不会再执行
}
代码执行的结果：

外层开始
外层即将异常
内层开始
内层异常前recover后的defer
内层准备recover
"内层"-"异常信息"
内层完成recover
内层recover前的defer
外层异常后
外层异常后defer
外层准备recover
外层没做啥事
外层完成recover
gibson1112
   gibson1112

  185***72536@163.com

1年前 (2019-05-23)
   朝阳群众

  cha***ngquynzhong@gmail.com

3
这个例子不给力，我重写了一个:

package main

import (
    "fmt"
)

// 自定义错误信息结构
type DIV_ERR struct {   
    etype int  // 错误类型   
    v1 int     // 记录下出错时的除数、被除数   
    v2 int
}
// 实现接口方法 error.Error()
func (div_err DIV_ERR) Error() string {   
    if 0==div_err.etype {      
        return "除零错误"   
    }else{   
        return "其他未知错误"  
    }
}
// 除法
func div(a int, b int) (int,*DIV_ERR) {  
    if b == 0 {     
        // 返回错误信息    
        return 0,&DIV_ERR{0,a,b}  
    } else {   
        // 返回正确的商  
        return a / b, nil   
    }
}
func main() { 
    // 正确调用  
    v,r :=div(100,2)  
    if nil!=r{   
        fmt.Println("(1)fail:",r)  
    }else{   
        fmt.Println("(1)succeed:",v) 
    }   
    // 错误调用
    v,r =div(100,0) 
    if nil!=r{   
        fmt.Println("(2)fail:",r)  
    }else{  
        fmt.Println("(2)succeed:",v) 
    }
}