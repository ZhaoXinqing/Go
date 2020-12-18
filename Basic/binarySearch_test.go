package Basic

import (
	"reflect"
	"testing"
)

//func Test_binarySearch(t *testing.T) {
//	arr := []int{4, 5, 6, 7, 8, 10, 13, 14, 15, 16, 17, 18, 19, 21}
//	a := 16
//	expectInt := 3
//	type args struct {
//		arr []int
//		a   int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{name: "binarySearch", args: args{arr: arr,a: a}, want: expectInt},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := binarySearch(tt.args.arr, tt.args.a); got != tt.want {
//				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_BubbleSort(t *testing.T) {
	arr := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}
	expectArr := []int{30, 18, 25, 36, 14, 17, 12, 9, 5, 7, 10}

	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "bubbleSort", args: args{arr: arr}, want: expectArr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
