// 函数将变量作为参数：Function1(recv)
// 方法在变量上被调用：recv.Method1()
// 接收者必须有一个显式的名字，这个名字必须在方法中被使用。
// receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

package main

import (
	"fmt"
)


// func (a int) Add (b int){    //方法非法！不能是内置数据类型
//   fmt.Println(a+b)
// }

// 合法的方法定义如下：

type myInt int

// Add ...
func Add(a, b int) { //函数
	fmt.Println(a + b)
}

func (a myInt) Add(b myInt) { //方法
	fmt.Println(a + b)
}

func main() {
	a, b := 3, 4
	var aa, bb myInt = 3, 4
	Add(a, b)
	aa.Add(bb)
}

// 上面的表达式aa.Add称作选择子（selector），它为接收者aa选择合适的Add方法。
