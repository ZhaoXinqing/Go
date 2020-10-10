func (recevier type) methodName(parameter list)(Return list){
	// 返回值和参数可以省略
}

package temp3

type Test struct{}

func (t Test) method0() {
	// 
}

package main

import (
	"fmt"
)

// 结构体
type User struct {
	Name string
	Email string
}

// 方法
func (u User) Notify() {
	fmt.Printf("%v : %v \n",u.Name,u.Email)
}
func main() {
	// 值类型调用方法
	u1 := User{"golang","golang@golang.com"}
	u1.Notify()
	// 指针类型调用方法
	u2 := User{"go","go@go.com"}
	u3 := &u2
	u3.Notify()
}

// 1.函数
func valueIntTest(a int) int {
	return a + 10
}

func pointIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:",valueIntTest(a))

	b := 5
	fmt.Println("pointerInTest:",pointerInTest(&b))
}

// 2.方法
type PersonD struct {
	id interface
	name string
}

// 接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}
func(p *PersonD) pointerShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	// 值类型调用方法
	personPointer := PersonD{101,"hello world"}
	personValue.valueShowName()
}

deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse

iris文档阅读学习
goadmin,iris文档阅读了解，
文档书写：环境部署及项目运行和测试。