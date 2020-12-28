package basic

import "testing"

func TestBucketSort(t *testing.T) {
	array := []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	num := 3
	expectArray := []int{1, 5, 7, 10, 12, 17, 18, 18, 23, 30}
	type args struct {
		theArray []int
		num      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestBucketSort", args: args{theArray: array, num: num}, want: expectArray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
