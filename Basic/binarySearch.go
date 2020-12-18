package Basic

//// 假设数组是从小到大排序
//func binarySearch(arr []int, a int) int {
//	var low int = 0
//	var height int = len(arr) - 1
//	for low <= height {
//		var mid int = low + (height-low)/2
//		var midValue int = arr[mid]
//		if midValue == a {
//			return mid
//		} else if midValue > a {
//			height = mid - 1
//		} else if midValue < a {
//			low = mid + 1
//		}
//	}
//	return -1
//}

// 冒泡排序
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
