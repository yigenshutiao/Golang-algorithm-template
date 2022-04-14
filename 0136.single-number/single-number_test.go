package _136_single_number

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {

	fmt.Println(10 ^ 0)
	fmt.Println(10 ^ 10)

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{2, 2, 1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
