// 结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
// 	Title ：标题
// 	Author ： 作者
// 	Subject：学科
// 	ID：书籍ID

// Go 语言中数组可以存储同一类型的数据，
// 但在结构体中我们可以为不同项定义不同的数据类型。

// 定义结构体
// 需要使用 type语句（设定结构体的名称） 和 struct语句（定义一个新的数据类型）；
type struct_variable_type struct {
	member definition
	member definition
	...
	member definition
 }

//  结构体变量声明
variable_name := structure_variable_type {value1, value2...valuen}
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

// 实例
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}) 
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

// 结果为：
// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Go 语言 www.runoob.com  0}