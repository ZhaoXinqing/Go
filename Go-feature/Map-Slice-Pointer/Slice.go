// 声明
var Slice_name []type

// 文邹邹的声明方式
name := make([]byte,length,cpacity)
// capacity指的是此切片当前指向内存的数据大小。而length指的是当前切片的容量大小
// 用切片操作目标内存的时候，必须小心别append过头，否则就操作到新开启的内存块去了，
// 也要小心别意外覆盖了原slice的值

// 判断切片是否为空,使用nil
numbers == nil
// true表示空切片，false表示非空切片

// 切片同数组相比，最灵活的方面在于切分子切片
numbers := []int{0,1,2,3,4,5,6,7,8}
number2 := numbers[:2] // 从0到2，但不包括2.所以是0，1
number3 := numbers[2:5] // 从2到5，但不包括5.所以是2，3，4
number4 := numbers[5:] // 从5到末尾，包括末尾。
// 子切片都是指向了源切片某一块内存，源切片元素发生了变化，那么子切片也会发生变化

// append()需要赋值给源切片，不然会值相同
s := []int{10} //创建一个legnth = capacity = 1 的切片，并且初始化为10
s = append(s,11) //容量不够，翻倍扩容。legnth = capacity = 2，现在是10，11
s = append(s,12) //容量又不够了，再次扩容。legnth =3， capacity = 4，现在是10，11，12
x := append(s, 13) //容量够了，不扩容。legnth = capacity = 4，现在是10，11，12，13
// 坑来了
y := append(s, 14) //容量够了，不扩容。legnth = capacity = 4，现在是10，11，12，14

// 不想受源切片影响，使用copy()函数。重新创建一个切片，自立山头
number5 := make([]int,2)
copy(number5,numbers[:2])


// 切片的数据结构：
type sLice struct {
	array unsafe.Pointer
	len int
	cap int
}
// 遇到切片时刻记住，用源切片来接受返回值。如果需要子切片，首要需要考虑是不是需要用copy来复制生成。

// 迭代数组和切片
numbers := []int{0,1,2,3,4,5,6,7,8}
for i := range numbers {
	fmt.Println("Slice item",1,"is",numbers[i])
}