package main

import (
	"errors"
	"fmt"
)

// 链表实现队列的方法定义
type node struct {
	Item interface{}
	Next *node
}

type linkedListQueue struct {
	Length int
	head   *node //头节点
	tail   *node //尾节点
}

func NewNode() *node {
	return &node{
		Item: nil,
		Next: nil,
	}
}

func NewLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (l *linkedListQueue) IsEmpty() bool {
	return l.Length == 0
}

func (l *linkedListQueue) Len() int {
	return l.Length
}

func (l *linkedListQueue) Enqueue(item interface{}) {
	buf := &node{
		Item: item,
		Next: nil,
	}
	if l.Length == 0 {
		l.tail = buf
		l.head = buf

	} else {
		l.tail.Next = buf
		l.tail = l.tail.Next
	}
	l.Length++
}

func (l *linkedListQueue) Dequeue() (item interface{}) {
	if l.Length == 0 {
		return errors.New(
			"failed to dequeue, queue is empty")
	}

	item = l.head.Item
	l.head = l.head.Next

	// 当只有一个元素时，出列后head和tail都等于nil
	// 这时要将tail置为nil，不然tail还会指向第一个元素的位置
	// 比如唯一的元素原本为2，不做这步tail还会指向2
	if l.Length == 1 {
		l.tail = nil
	}

	l.Length--
	return
}

func (l *linkedListQueue) Traverse() (resp []interface{}) {
	buf := l.head
	for i := 0; i < l.Length; i++ {
		resp = append(resp, buf.Item, "<--")
		buf = buf.Next
	}
	return
}

func (l *linkedListQueue) GetHead() (item interface{}) {
	if l.Length == 0 {
		return errors.New(
			"failed to getHead, queue is empty")
	}
	return l.head.Item
}

// 测试结果
func main() {
	q := linked_list3.NewLinkedListQueue()

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
		fmt.Println(q.Traverse())
	}

	fmt.Println(q.GetHead(), q.Len(), q.IsEmpty())

	for i := 0; i < 5; i++ {
		q.Dequeue()
		fmt.Println(q.Traverse())
	}
}

// // 输出结果
// [0 <--]
// [0 <-- 1 <--]
// [0 <-- 1 <-- 2 <--]
// [0 <-- 1 <-- 2 <-- 3 <--]
// [0 <-- 1 <-- 2 <-- 3 <-- 4 <--]
// 0 5 false
// [1 <-- 2 <-- 3 <-- 4 <--]
// [2 <-- 3 <-- 4 <--]
// [3 <-- 4 <--]
// [4 <--]
// []

// 作者：哥斯拉啊啊啊哦
// 链接：https://www.jianshu.com/p/6cb6798ecef0
// 来源：简书
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
