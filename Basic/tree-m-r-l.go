package Basic

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func trans(root *Student) {
	// 前序遍历：先遍历根节点再遍历左子树，再遍历右子树
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func trans1(root *Student) {
	// 中序遍历：先遍历左子树再遍历跟，再遍历根节点再遍历右子树
	if root == nil {
		return
	}
	trans1(root.left)
	fmt.Println(root)
	trans1(root.right)
}

func trans2(root *Student) {
	// 后序遍历：先遍历左子树再遍历跟，再遍历右子树，再遍历根节点
	if root == nil {
		return
	}
	trans1(root.left)
	trans1(root.right)
	fmt.Println(root)
}

// 备注，所谓的前序，中序，后序，都是针对根节点而言的

func main() {
	var root *Student = new(Student)
	root.Name = "stu01"
	root.Age = 18
	root.Score = 100
	var left1 *Student = new(Student)
	left1.Name = "stu02"
	left1.Age = 18
	left1.Score = 100
	root.left = left1
	var right1 *Student = new(Student)
	right1.Name = "stu04"
	right1.Age = 18
	right1.Score = 200
	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "stu03"
	left2.Age = 18
	left2.Score = 180
	left1.left = left2
	//遍历
	fmt.Println("先序遍历：")
	trans(root)
	fmt.Println("中序遍历：")
	trans1(root)
	fmt.Println("后序遍历：")
	trans2(root)
}
