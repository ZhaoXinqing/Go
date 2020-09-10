package main

import (
	"container/list"
	"fmt"
)

// Binary Tree
type BinaryTree struct {
	Data interface{}
	Left *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data interface{}) *BinaryTree {
	
}