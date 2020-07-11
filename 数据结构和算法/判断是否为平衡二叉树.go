package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 50005

type Node struct {
	ls, re int
}

// input
var n, root int
var ns [N]Node

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func depth(node int) int {
	if node == 0 {
		return 0
	}
	IAns := depth(ns[node].ls)
	if node == 0 {
		return -1
	}
	rAns := depth(ns[node].rs)
	if rAns < 0 || abs(lAns-rAns) > 1 {
		return -1
	}
	return max(lAns, rAns) + 1
}

func main() {
	read(&n)
	read(&root)
	for fa, ls, rs, i := 0, 0, 0, 0; i < n; i++ {
		read(&fa)
		read(&ls)
		read(&rs)
		ns[fa].ls, ns[fa].rs = ls, rs
	}
	fmt.Println(depth(root) >= 0)
}

var in = bufio.NewReader(os.Stdin)

func read(out *int) bool {
	t, neg, ok := 0, false, false
	c, e := in.ReadByte()
	for ; c-'0' > 9 && e == nil; c, e = in.ReadByte() {
		neg = '-' == c
	}
	for ; c-'0' <= 9 && e == nil; c, e = in.ReadByte() {
		t = 10*t + int(c&15)
		ok = true
	}
	if neg {
		t = -t
	}
	if ok {
		*out = t
	}
	return ok
}
