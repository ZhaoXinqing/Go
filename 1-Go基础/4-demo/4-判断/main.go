// 判断二叉树是否为满二叉树
// 满二叉树的定义：一个高度为h，并且含有2^h - 1个节点的二叉树称为满二叉树，下文称呼满二叉树为FBT。
// 根据满二叉树的高度与节点个数之间的关系，很容易判断一棵树是否为FBT，只需要求树其树高和节点个数即可

package main

import "container/list"

//思路:层序遍历,只需分两种情况
//1）若当前节点左孩子为空右孩子不为空,直接返回不空
//2）若当前节点的左右孩子不全,则其后面的节点必须都为叶节点
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := list.New()
	l.PushBack(root)
	leaf := false //第二种情况,状态开启
	for l.Len() > 0 {
		node := l.Front().Value.(*TreeNode)
		left := node.Left //得到左右孩子进行判断
		right := node.Right
		//第一个判断:如果已经开启leaf状态,后面不是页节点的直接返回false
		//第二个判断,左孩子为空而右孩子不为空
		if leaf && (left != nil || right != nil) || left == nil && right != nil {
			return false
		}
		//继续判断
		if left != nil {
			l.PushBack(left)
		}
		if right != nil {
			l.PushBack(right)
		}
		//左右孩子不全,开启叶子节点状态为true
		if left == nil || right == nil {
			leaf = true
		}
	}
	return true
}
