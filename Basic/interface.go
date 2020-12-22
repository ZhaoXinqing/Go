package basic

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// interfaceToString 接口转string
func interfaceToString(inter interface{}) {
	switch inter.(type) {
	case string:
		fmt.Println("string", inter.(string))
		break
	case int:
		fmt.Println("int", inter.(int))
		break
	case float64:
		fmt.Println("float64", inter.(float64))
		break
	}
}

// sliceToMap 切片转换为Map
func sliceToMap(s []string) map[int]string {
	mapObj := map[int]string{}
	for key, value := range s {
		mapObj[key] = value
	}
	return mapObj
}

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// reverseSlice 切片反序
func reverseSlice(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// reverseWords 单词反转
func reverseWords(s string) string {
	es := strings.Split(s, " ")
	es1 := make([]string, 0, len(es))
	for i := len(es) - 1; i >= 0; i-- {
		if 0 != len(es[i]) {
			es1 = append(es1, es[i])
		}
	}
	return strings.Join(es1, " ")
}

//快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]
	low := make([]int, 0, 0)
	height := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			height = append(height, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, height = QuickSort(low), QuickSort(height)
	myArr := append(append(low, mid...), height...)
	return myArr
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

func binarySearch02(arr []int, a int) int {
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

//RemoveRepByMap slice去重
func RemoveRepByMap(arr []int64) []int64 {
	result := []int64{1}        //存放返回的不重复切片
	tempMap := map[int64]byte{} // 存放不重复主键
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

///***
// *
// *                 .-~~~~~~~~~-._       _.-~~~~~~~~~-.
// *             __.'  欢迎访问    ~.   .~              `.__
// *           .'//  我的个人博客    \./   ☞ GitHub.com ☜   \\`.
// *         .'//                     |                     \\`.
// *       .'// .-~"""""""~~~~-._     |     _,-~~~~"""""""~-. \\`.
// *     .'//.-"                 `-.  |  .-'                 "-.\\`.
// *   .'//______.============-..   \ | /   ..-============.______\\`.
// * .'______________________________\|/______________________________`.
// */

//
// 插入排序
func insertSort(slice []int) {
	if len(slice) <= 0 {
		return
	}
	for i := 1; i < len(slice); i++ {
		temp := slice[i]
		for j := i - 1; j >= 0; j-- {
			if temp < slice[j] {
				slice[j+1] = slice[j]
			}
			slice[j+1] = temp
			break
		}
	}
}

// 堆排序
// 1-取出当前队列的最大值
func heapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		topmax := i
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

// 2-排序
func heapSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < len(arr); i++ {
		lastLen := length - i
		heapSortMax(arr, lastLen) // 从小到大的排序
		// 调成从大到小
		if i < length {
			arr[0], arr[lastLen-1] = arr[length-1], arr[0]
		}
	}
	return arr
}

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

// RemoveDuplicate 整型数组去重
func RemoveDuplicate(arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	tempMap := map[int64]byte{}
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// insertSort 插入排序
func insertSort02(slice []int) []int { // 没有返回值，可直接调用
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
	return slice
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

// FindDir 遍历文件夹
func FindDir(dir string, num int) interface{} {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var a interface{}
	for _, fi := range fileInfo {
		a, _ = fmt.Printf(strings.Repeat("\t", num))
		if fi.IsDir() {
			a, _ = fmt.Printf(`目录：`, fi.Name())
			FindDir(dir+`/`+fi.Name(), num+1)
		} else {
			a, _ = fmt.Printf(`文件：`, fi.Name())
		}
	}
	return a
}
