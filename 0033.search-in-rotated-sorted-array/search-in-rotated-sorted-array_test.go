package _033_search_in_rotated_sorted_array

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"name2",
			args{
				[]int{4, 5, 6, 7, 0, 1, 2},
				0,
			},
			4,
		},

		{
			"test 1",
			args{
				[]int{3, 1},
				1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
