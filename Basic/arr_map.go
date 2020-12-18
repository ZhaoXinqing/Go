package Basic

//
//func quickSort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//	splitdata := arr[0]
//	low := make([]int, 0, 0)
//	hight := make([]int, 0, 0)
//	mid := make([]int, 0, 0)
//	mid = append(mid, splitdata)
//	for i := 1; i < len(arr); i++ {
//		if arr[i] < splitdata {
//			low = append(low, arr[i])
//		} else if arr[i] > splitdata {
//			hight = append(hight, arr[i])
//		} else {
//			mid = append(mid, arr[i])
//		}
//	}
//	low, hight = quickSort(low), quickSort(hight)
//	myarr := append(append(low, mid...), hight...)
//	return myarr
//}
//
//// 堆排序
//// 1-取出当前队列的最大值
//func heapSortMax(arr []int, length int) []int {
//	if length <= 1 {
//		return arr
//	}
//	depth := length/2 - 1
//	for i := depth; i >= 0; i-- {
//		topmax := i
//		leftchild := 2*i + 1
//		rightchild := 2*i + 2
//		if leftchild <= length && arr[leftchild] > arr[topmax] {
//			topmax = leftchild
//		}
//		if rightchild <= length && arr[rightchild] > arr[topmax] {
//			topmax = rightchild
//		}
//		if topmax != i {
//			arr[i], arr[topmax] = arr[topmax], arr[i]
//		}
//	}
//	return arr
//}
//
//// 2-排序
//func heapSort(arr []int) []int {
//	length := len(arr)
//	for i := 1; i < len(arr); i++ {
//		lastlen := length - i
//		heapSortMax(arr, lastlen)  // 从小到大的排序
//		// 调成从大到小
//		if i < lenth {
//			arr[0],arr[lastlen-1]= arr[lenlen-1],arr[0]
//		}
//		return arr
//	}
//}

//func main() {
//	arr := []int{4, 5, 6, 7, 8, 10, 13, 14, 15, 16, 17, 18, 19, 21}
//	b := binarySearch(arr, 16)
//	if b >= 0 {
//		fmt.Println("目标值下标围是：", b)
//	} else {
//		fmt.Println("查找目标不存在")
//	}
//}
//
//// 时间复杂度：log(n)
//
//// 插入排序
//func insertSort(slice []int) {
//	if len(slice) <= 0 {
//		return
//	}
//	for i := 1; i < len(slice); i++ {
//		temp := slice[i]
//		for j := i - 1; j >= 0; j-- {
//			if temp < slice[j] {
//				slice[j+1] = slice[j]
//			}
//			slice[j+1] = temp
//			break
//		}
//	}
//}
//
