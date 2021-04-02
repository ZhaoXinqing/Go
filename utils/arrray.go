package utils

// IsContain ...
func IsContain(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsContainStr ...
func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
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
