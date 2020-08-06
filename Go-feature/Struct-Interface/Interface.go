// 接口 是 方法特征 的命名集合
// 在Go语言的实现中，一个interface类型的变量存储了2个信息, 一个《值，类型》对(value, type)
// value是实际变量值，type是实际变量的类型。

package Interface

import "fmt"

// Phone 定义接口
type Phone interface {
	call(param int) string
	takephoto()
}

// Huawei 定义结构体
type Huawei struct {
}

// 定义方法 call - 带参数的方法
func (huawei Huawei) call(param int) string {
	fmt.Println("i am Huawei, i can call you!", param)
	return "damon"
}

// 定义方法 takephoto
func (huawei Huawei) takephoto() {
	fmt.Println("i can take a photo for you")
}

func main() {
	var phone Phone     // 定义接口Phone 接口变量 phone
	phone = new(Huawei) // 声明变量
	phone.takephoto()   // 实现接口变量方法
	r := phone.call(50) //
	fmt.Println(r)
}
