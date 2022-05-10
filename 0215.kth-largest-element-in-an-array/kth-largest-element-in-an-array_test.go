package _215_kth_largest_element_in_an_array

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"test 1",
		//	args{
		//		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
		//		1,
		//	},
		//	11,
		//},
		{
			"test 1",
			args{
				[]int{6, 5, 4, 3, 2, 1},
				2,
			},
			5,
		},
		{
			"test 2",
			args{
				[]int{2, 1},
				2,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
