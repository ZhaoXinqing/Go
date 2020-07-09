// 单线程
// package main

// import (
//    "fmt"
//   )

// func main(){
//     array := []int{3, 6, 1, 4, 2, 8}
//     fmt.Println(array)
//     quickSort(array, 0, len(array)-1)
//     fmt.Println(array)
// }

// func quickSort(array []int, left, right int) {
//     if left >= right {
//         return
//     }
//     index := partition(array,left,right)
//     quickSort(array, left, index - 1)
//     quickSort(array, index + 1, right)
// }

// func partition(array []int, left, right int) int {
//     baseNum := array[left]
//     for left < right{
//         for (array[right] >= baseNum && right > left){
//             right--
//         }
//         array[left] = array[right]
//         for (array[left] <=baseNum && right > left) {
//             left++
//         }
//         array[right] = array[left]
//     }
//     array[right] = baseNum
//     return right
// }

// 多线程
package main

import (
	"fmt"
)

func main() {
	array := []int{3, 1, 4, 1, 5, 9, 2, 6}
	ch := make(chan int)
	go quickSort(array, ch)
	for value := range ch {
		fmt.Println(value)
	}
}
func quickSort(array []int, ch chan int) {
	if len(array) == 1 {
		ch <- array[0]
		close(ch)
		return
	}
	if len(array) == 0 {
		close(ch)
		return
	}
	small := make([]int, 0)
	big := make([]int, 0)
	left := array[0]
	array = array[1:]
	for _, num := range array {
		switch {
		case num <= left:
			small = append(small, num)
		case num > left:
			big = append(big, num)
		}
	}
	leftCh := make(chan int, len(small))
	rightCh := make(chan int, len(big))

	go quickSort(small, leftCh)
	go quickSort(big, rightCh)

	//合并数据
	for i := range leftCh {
		ch <- i
	}
	ch <- left
	for i := range rightCh {
		ch <- i
	}
	close(ch)
	return
}
