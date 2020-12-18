package Basic

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestRemoveRepByMap", args: args{a: i}, want: a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
