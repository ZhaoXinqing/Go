// 类型操作有两种：一种是类型转换，一种是类型断言
// Golang的类型分为两种，一种是静态类型，一种是底层类型

type myString string
// 通过 type 关键字声明了自己的类型 myString ，其中 myString 是静态类型，string 则是底层类型；
// 如果底层类型相同，则两种类型可以转换成功，反之则失败；
// 如果不知道底层类型，则可以调用reflect的Kind方法来获取：
fmt.Printf("%v",reflect.Kind(count))

// 类型断言：
// 很多时候，会将普通类型转换成interface{}，这个过程是隐式推导；
// 但如果反向，将interface{}向普通类型进行转换，则需要显式推导，
// 这个时候就需要类型断言：共有两种转换方式：

// 第一种：Comma-ok: value, ok := element.(T)
var a intreface{}
value,ok := a.(string)
if != ok {
	fmt.Println("it is not ok for type string")
	return
}

// 第二种：switch判断：(优点是节省代码，本质和Comma-OK一样)
var t interface{}
t = functionOfSomeType()

switch t := t.(type) {
case bool:
	fmt.Println("boolean %t\n",t)
case int:
	fmt.Println("interger %d\n",t)
case *bool:
	fmt.Println("pointer to boolean %t\t",*t)
case *int:
	fmt.Println("pointer to integer %d\n",*t)
default:
	fmt.Printf("unexpected type %T",t)
}


