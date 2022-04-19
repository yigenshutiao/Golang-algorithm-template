package _031_next_permutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				[]int{1, 2, 3, 8, 5, 7, 6, 4},
			},
			[]int{1, 2, 3, 8, 6, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
