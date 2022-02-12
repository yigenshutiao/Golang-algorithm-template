package _128_longest_consecutive_sequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			"test 1",
			args{
				[]int{100, 4, 200, 1, 3, 2},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := longestConsecutive(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("longestConsecutive() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
