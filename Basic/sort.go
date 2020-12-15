package main

import (
	"fmt"
)

// insertSort 插入排序
func insertSort(slice []int) { // 没有返回值，可直接调用
	if len(slice) <= 0 {
		return
	}
	for i := 1; i < len(slice); i++ {
		//下标i的元素temp
		temp := slice[i]
		for j := i - 1; j >= 0; j-- {
			//从i-1的位置往前查找，如果前边的元素比temp大就进行后移
			if temp < slice[j] {
				slice[j+1] = slice[j]
			}
			//否则找到比temp小的就将tmep插入。
			slice[j+1] = temp
			break
		}
	}
}

// BubbleSort 冒泡排序，稳定
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// QuickSort 快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

// HeapSortMax 每次从堆中 取一个最大数 放到 i处
func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length && arr[leftchild] > arr[topmax] {
			topmax = leftchild
		}
		if rightchild <= length && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}

// HeapSort 堆排序
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}

// shellSort 希尔排序
func shellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap > 0 {
		for i := 0; i < len(arr); i++ {
			j := i
			for (j-gap >= 0) && (arr[j] < arr[j-gap]) {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
		gap /= 2
		fmt.Println(arr)
	}
	return arr
}

// mergerSort 归并排序
/**
归并排序（Merge sort，台湾译作：合并排序）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用
分治思想，时间复杂度为：O(n*log(n))
*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...) // 这里竟然没有报数组越界的异常？
	result = append(result, left[m:]...)
	return result
}

func main() {
	a := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	b := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	c := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	d := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	e := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	f := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	insertSort(a)                         // 插入排序（只能小规模数据，不然效率低）
	fmt.Println("冒泡排序：b", BubbleSort(b))  // 冒泡排序
	fmt.Println("快速排序：c", QuickSort(c))   // 快速排序
	fmt.Println("堆排序：d", HeapSort(d, 11)) // 堆排序
	fmt.Println("希尔排序：e", shellSort(e))   // 希尔排序（改进的插入排序，可以对应大规模数据）
	fmt.Println("归并排序：f", mergeSort(f))   // 归并排序
}



func binarySearch(arr []int,a int)int {
	var low int = 0
	var height int = len(arr) - 1
	for low <= height {
		var mid int = low + (height-low) /2
		var midValue int = arr[mid]
		if midValue == a {
			return mid
		} else if midValue > a {
			height = mid -1
		} else if midValue < a{
			low = mid + 1
		}
	}
	return - 1
}

a := []int{1,2,3,4,5,7,8,9,12}
b := binarySearch(a, 5)
if b >= 0 {
	fmt.Println("目标下标值是：", b)
} else {
	fmt.Println("查找目标不存在")
}
