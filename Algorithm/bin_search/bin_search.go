// 有序的队列中。将数组一分为二
// 相等直接返回
// 元素大于分割点，在分割点右侧继续查找
// 元素小于分割点，在分割点左侧继续查找

// 时间复杂度：O（lgn）

// 有序的数组，并能支持随机访问

// 用户ip区间段查询
// 相似度查询

package main

import "fmt"

func bin_search(arr []int, finddata int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] > finddata {
			high = mid - 1
		} else if arr[mid] < finddata {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := bin_search(arr, 1024)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}
