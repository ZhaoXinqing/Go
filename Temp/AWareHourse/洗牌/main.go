package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //设置种子

	sixah := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	rand.Shuffle(len(sixah), func(i, j int) { //调用算法
		sixah[i], sixah[j] = sixah[j], sixah[i]
	})

	log.Println(sixah)
	return
}

// 高校洗牌
// https://studygolang.com/articles/24498?fr=sidebar