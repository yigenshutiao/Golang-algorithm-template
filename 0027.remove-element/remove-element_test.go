package _027_remove_element

import "testing"

func Test_removeElements(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{0, 1, 2, 2, 3, 0, 4, 2},
				2,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
