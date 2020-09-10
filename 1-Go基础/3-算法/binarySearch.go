package main

import (
	"fmt"
)

// 假设数组是从小到大排序
func binarySearch(arr []int, a int) int {
	var low int = 0
	var height int = len(arr) - 1
	for low <= height {
		var mid int = low + (height-low)/2
		var midValue int = arr[mid]
		if midValue == a {
			return mid
		} else if midValue > a {
			height = mid - 1
		} else if midValue < a {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 8, 10, 13, 14, 15, 16, 17, 18, 19, 21}
	b := binarySearch(arr, 16)
	if b >= 0 {
		fmt.Println("目标值下标围是：", b)
	} else {
		fmt.Println("查找目标不存在")
	}
}

// 时间复杂度：log(n)
