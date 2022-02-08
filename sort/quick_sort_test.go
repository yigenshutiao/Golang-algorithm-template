package main

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test 1",
			args{
				[]int{2, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
