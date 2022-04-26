package _283_move_zeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes21(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[]int{0, 1, 0, 3, 12},
			},
			[]int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[]int{1, 0, 0, 3, 12},
			},
			[]int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
