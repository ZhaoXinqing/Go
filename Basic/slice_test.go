package Basic

import (
	"reflect"
	"testing"
)

func TestRemoveRepByMap(t *testing.T) {
	arr := []int64{12, 15, 20, 20, 26, 26, 28}
	expectArr := []int64{12, 15, 20, 26, 28}

	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{name: "TestRemoveRepByMap", args: args{arr: arr}, want: expectArr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRepByMap(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRepByMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
