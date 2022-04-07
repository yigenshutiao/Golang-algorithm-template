package _042_trapping_rain_water

import "testing"

func Test_traps(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{4, 2, 0, 3, 2, 5},
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("traps() = %v, want %v", got, tt.want)
			}
		})
	}
}
