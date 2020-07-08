package user

import (
	"strconv"
)

type User struct {
	name  string
	age   int
	sex   string
	phone string
}

func (s *User) SetName(name string) {
	s.name = name
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) SetAge(age int) {
	s.age = age
}

func (s *User) GetAge() int {
	return s.age
}

func (s *User) String() string {
	return "name is " + s.name + ",age is " + strconv.Itoa(s.age) + " ,sex=" + s.sex + " ,phone=" + s.phone
}

func (s *User) SetSex(sex string) {
	s.sex = sex
}

func (s *User) GetSex() string {
	return s.sex
}

func (s *User) SetPhone(phone string) {
	s.phone = phone
}

func (s *User) GetPhone() string {
	return s.phone
}

func (User) Print() string {
	return "print"
}


定义方法
格式：
func (变量名 类型) 方法名(形参列表) (返回值列表) {
	#code
}

示例：
package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Test1(newName string) string {
	p.Name = newName
	return p.Name
}

func (p *Person) Test2(newName string) string {
	p.Name = newName
	return p.Name
}

func main() {
	p1 := Person{"abc"}
	tmpName := p1.Test1("xyz")
	fmt.Println(tmpName, p1.Name) //xyz  abc

	p2 := Person{"aaaa"}
	realName := p2.Test2("zzzzz")
	fmt.Println(realName, p2.Name) //zzzz zzzz
}
　　是否使用指针设计是否要修改调用方法的结构体内部成员属性，调用的时候，可以不用关心定义方法时是使用指针还是变量。

　　如果是非指针类型，那么就是值传递，在方法内部修改成员属性，并不会反映到实际的结构体上，除非使用指针形式。

　　

方法表达式
package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Test1(newName string) string {
	p.Name = newName
	return p.Name
}

func (p *Person) Test2(newName string) string {
	p.Name = newName
	return p.Name
}

func main() {
	Pfunc1 := Person.Test1
	Pfunc2 := (*Person).Test2

	p1 := Person{"abx"}
	tmp := Pfunc1(p1, "xxx")
	fmt.Println(tmp, p1.Name) //xyz  abx

	p2 := Person{"xyz"}
	realname := Pfunc2(&p2, "xxx")
	fmt.Println(realname, p2.Name) //xyz  xxx

}