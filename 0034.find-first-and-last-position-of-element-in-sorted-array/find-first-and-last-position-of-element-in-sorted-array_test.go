package _034_find_first_and_last_position_of_element_in_sorted_array

import "testing"

func Test_rightSearch(t *testing.T) {
	type args struct {
		num    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{5, 7},
				7,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSearch(tt.args.num, tt.args.target); got != tt.want {
				t.Errorf("rightSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
