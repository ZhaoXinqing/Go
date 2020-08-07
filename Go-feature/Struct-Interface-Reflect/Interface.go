// 声明一个接口：
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	...
	method_namen [return_type]
}
// 声明struct时候，是type name struct，在struct里面，都是各个成员属性
// 声明接口，是用type name interface，在interface里面，就都是各个成员函数
// 使用示例:
type Shape interface {
	area() float64
}
func getArea(shape Shape) float64{
	return shape.area()
}
// 声明一个Shape的interface,里面有一个成员函数，是area，返回一个float64；
// 而getArea只接受Shape接口类型

// 实现类
type Circle struct {
	x,y,radius float64
}
func(circle Circle)area() float64{
	return math.Pi * circle.radius * circle.radius
}
// 现在Circle类就实现了刚才Shape的接口。验证一下:
circle := Circle{x:0,y:0,radius:5}
fmt.Println("Circle area: %f\n",getArea(circle))

// 只有一个实现类，那就体现不出接口的优势了。再上一个实现类:
type Rectangle struct {
	width,height float64
}
func(rect Rectangle) area() float64 {
	return rect.width * rect.height
}
rectangle := Rectangle {width:10,height:5}
fmt.Println("Rectangle area:%f\n",getArea(rectangle))
// 现在Rectangle和Circle都实现了Shape接口类。
// 实现的秘密就在于Rectangle和Circle两个类都有area() float64这个函数

// Golang中接口的灵活性可以大大减少代码量，减少耦合性。但也反过来降低了可读性
// 建议在写代码的时候，一定要记得输出日志，最好能在关键节点输出尽可能的详细日志。

// 在Golang当中，如何实现一个接口类？那就是把接口类中的函数都实现了，就成。

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
