package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int64{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	expectArray := []int64{1, 5, 7, 10, 12, 17, 18, 18, 23, 30}
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "QuickSort", args: args{arr: array}, want: expectArray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println(expectArray)
}
