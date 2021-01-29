## 指针

在 Go 中 * 代表取指针地址中存的值，& 代表取一个值的地址
定义一个指针并像普通变量那样给他赋值，代码如下；
	var i *int
	*i = 1
	fmt.Println(i, &i, *i)

//就会报这样的一个错误
//panic: runtime error: invalid memory address or nil pointer dereference
//[signal 0xc0000005 code=0x1 addr=0x0 pc=0x498025]

//报这个错的原因是 go 初始化指针的时候会为指针 i 的值赋为 nil ，但 i 的值代表的是 *i 的地址， nil的话系统还并没有给 *i 分配地址，所以这时给 *i 赋值肯定会出错
//解决这个问题非常简单，在给指针赋值前可以先创建一块内存分配给赋值对象即可：
   var i *int
   i = new(int)
   *i = 1
   fmt.Println(i, &i, *i)
