package main

import (
	"fmt"
	"user"
)

type Person struct {
	ID      string
	Name    string
	Address string
}

func main() {
	userDB := make(map[string](*user.User)) //注意写法

	//初始化，注意对数组的初始化

	u := new(user.User)
	u.SetAge(12)
	u.SetName("张三")
	u.SetSex("男")
	u.SetPhone("15902783102")

	userDB["u1"] = u

	v, ok := userDB["u1"] //获取map值  key为'u1'
	if !ok {
		fmt.Println(" 没有找到信息")
		return
	}

	//打印出全部值 和各个分值
	fmt.Println(v.String())
	fmt.Printf("userDB[u1] = \n {name=%v \n age=%v \n sex=%v \n phone=%v \n}", v.GetName(), v.GetAge(), v.GetSex(), v.GetPhone())

	personDB := make(map[string]Person)
	var person = Person{"12346", "Xym", "bbb"}
	personDB["p1"] = person
	p, ok := personDB["p1"] //获取map值  key为'u1'
	if !ok {
		fmt.Println(" 没有找到信息")
		return
	}

	//打印出全部值 和各个分值
	fmt.Printf(" \n personDB[p1] = \n {ID=%v \n Name=%v \n address=%v \n}", p.ID, p.Name, p.Address)

	// 先声明map
	//var a map[string]string
	//a["b"] = "c" //这样会报错的，要先初始化内存
	//a = make(map[string]string)
	//a["b"] = "c" //这样才不会错

	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "cc"
	m1["d"] = "dd"
	fmt.Println("\n", m1) //输出元素的顺序是随机的，go语言中map是无序的

	// 直接创建
	m2 := make(map[string]string)
	m2["a"] = "aa"
	m2["b"] = "bb"
	fmt.Println("\n", m2)

	// 初始化 + 赋值一体化 注意最后一个元素后面也需要带上逗号的
	m3 := map[string]string{
		"a":  "aa",
		"b":  "bb",
		"c":  "cc",
		"xx": "xx",
		"d":  "dd", //注意最后一个元素后面也需要带上逗号的
	}
	fmt.Println("\n", m3)

	//删除一个key对应的元素 map的删除操作
	fmt.Println("删除之前", m3)
	delete(m3, "a")
	fmt.Println("删除之后", m3)

	//map的遍历
	fmt.Println("\n 遍历输出map")
	for k, v := range m3 {
		fmt.Printf("key=%v,val=%v \n", k, v) //无序的输出
	}
}
