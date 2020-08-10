// 类型操作有两种：一种是类型转换，一种是类型断言
// Golang的类型分为两种，一种是静态类型，一种是底层类型

type myString string
// 通过type关键字声明了自己的类型myString。myString则是静态类型，而string则是底层类型。
// 如果底层类型相同，则两种类型可以转换成功，反之则失败。
// 如果不知道底层类型，则可以调用reflect的Kind方法来获取，如下:
fmt.Printf("%v",reflect.Kind(count))


// 类型断言
// 在Golang当中还有接口类型interface{}。很多时候，会将普通类型转换成interface{}，
// 这个过程是隐式推导，用户无感知。但如果反向，将interface{}向普通类型进行转换，
// 则需要进行显示推导。这个时候就需要进行类型断言.

// 第一种，Comma-ok:
value,ok := element.(T)
// 示例
var a interface{}
value, ok := a.(string)
if !ok {
	fmt.Println("it is not ok for type string")
	return
}

//第二种，switch判断
// 示例
var t interface{}
t = functionOfSomeType()

switch t := t.(type) {
case bool:
	fmt.Println("boolean %t\n",t)
case int:
	fmt.Println("integer %d\n",t)
case *bool:
	fmt.Println("pointer to boolean %t\t",*t)
case *int:
	fmt.Println("pointer to integer %d\n",*t)
default:
	fmt.Printf("unexpected type %T",t)
}
// 优点就是可以节省代码，本质和Comma-OK一样

// 综合起来，golang类型转换就是三条规则:
		// 普通类型向interface{}转换是隐式。
		// interface{}向普通类型转换必须显示
		// 强制转换时，最好使用类型断言，防止panic。

	Golang在设计的时候，肯定是参考了目前市面上各种较为流行的语言，统筹考虑之后才变成今天这个样子的。取各家
		之所长，创自家之风骚。像Golang是大小写敏感的语言，既支持静态类型变量声明(var i int)又支持动态类型变
		量声明(i := 1)，还支持同时声明多个不同类型的变量(var a, b, c = 3,4,"foo")还有其它语法特点，天知道
		golang在设计的时候，汲取了多少营养。我们不是编译器大师也不是架构大师，与其吐槽golang的种种不好，不如
		细细分析golang能否提高自己的职业竞争力。