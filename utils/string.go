package utils

import "strconv"

// StrToFloat64 ...
func StrToFloat64(str string) (floatValue float64, err error) {
	floatValue, err = strconv.ParseFloat(str, 64)
	return
}

// StrToInt ...
func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// StrToInt64 ...
func StrToInt64(str string) (int64, error) {
	intNum, err := strconv.Atoi(str)
	int64Num := int64(intNum)
	return int64Num, err
}

// StrToUint64 ...
func StrToUint64(str string) (uint64, error) {
	intNum, err := strconv.Atoi(str)
	uint64Num := uint64(intNum)
	return uint64Num, err
}
