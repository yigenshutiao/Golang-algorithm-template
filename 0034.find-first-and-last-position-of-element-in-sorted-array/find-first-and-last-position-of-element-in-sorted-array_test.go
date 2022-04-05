package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			[]int{3, 4},
		},
		{
			"test 2",
			args{
				[]int{5, 7, 7, 8, 8},
				8,
			},
			[]int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
