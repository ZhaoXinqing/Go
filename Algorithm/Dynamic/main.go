package main

import (
	"fmt"
)

func SimpleGetNum(num int) int {
    var dp [10000]int
    moneys = int[...] {1, 5, 10,20, 50, 100}// 面额数组
    for i := 0; i <= num; i++ { // 用1rmb面额拼凑0金额的方法数都为1
        dp[i] = 1
    }
    
    for j := 1; j <= 5; j++ {
		for i := 1; i <= num; i++ {
			if i >= moneys[j] { // 当前面金额大于面额的时候需要计算前i种面额组合出i - moneys[j]的方法数
				dp[i] += dp[i - moneys[j]]
			}
		}
	}

	return dp[num]
}

func main() {
	fmt.Println(SimpleGetNum(30))
}
