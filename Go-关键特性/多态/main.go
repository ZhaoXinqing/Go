// 使用接口展示多态行为
package main

import (
	"fmt"
)

// notifier 是一个定义了通知类的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用指针接收者实现了notifier 接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin 定义了程序中的管理者
type admin struct {
	name  string
	email string
}

// notify使用指针接受者实现了 notifier 接口
func (a *admin) notify() {
	fmt.Printf("Send admin email to %s<%s\n", a.admin, a.email)
}

// main 是程序主入口
func main() {
	// 创建一个user值传给sendNotification
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// 创建一个 admin 值并传给sendNotification
	lisa := admin{"Lisa", "lias@email.com"}
	sendNotification(&lisa)
}

// senNotificaton 接受一个实现 notifier 接口的值
// 并发通知
func sendNotification(n notifier) {
	n.notify()
}
