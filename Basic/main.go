package Basic

//
//import (
//	"fmt"
//	"strconv"
//)
//
//// RemoveDuplicate 整型数组去重
//func RemoveDuplicate(arr []int64) []int64 {
//	result := make([]int64, 0, len(arr))
//	tempMap := map[int64]byte{}
//	for _, e := range arr {
//		l := len(tempMap)
//		tempMap[e] = 0
//		if len(tempMap) != l {
//			result = append(result, e)
//		}
//	}
//	return result
//}
////切片转换为Map
//func sliceToMap(s []string) map[int]string {
//	mapObj := map[int]string{}
//	for key,value := range s {
//		mapObj[key] = value
//	}
//	return mapObj
//}
//
//
//func main() {
//	//切片转换为Map
//	s := []string{"a", "b", "c", "d", "e"}
//	map := sliceToMap(s)
//	fmt.Println(map)
//
//
//	//map如何顺序读取?
//	s := []int64{10, 20, 10, 20, 30, 40}
//	s = RemoveRepByMap(s)
//	fmt.Println(s)
//
//
//	//排序
//	a := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	b := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	c := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	d := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	e := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	f := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
//	insertSort(a)                         // 插入排序（只能小规模数据，不然效率低）
//	fmt.Println("冒泡排序：b", BubbleSort(b))  // 冒泡排序
//	fmt.Println("快速排序：c", QuickSort(c))   // 快速排序
//	fmt.Println("堆排序：d", HeapSort(d, 11)) // 堆排序
//	fmt.Println("希尔排序：e", shellSort(e))   // 希尔排序（改进的插入排序，可以对应大规模数据）
//	fmt.Println("归并排序：f", mergeSort(f))   // 归并排序
//
//
//	// Go int、int64、string类型转换
//
//	// string 到 int64
//	var abc string = "101213"
//	int64, _ := strconv.ParseInt(abc, 10, 64)
//
//	var bcd int64
//	// int64 到 string
//	string := strconv.FormatInt(bcd, 10)
//
//
//}
