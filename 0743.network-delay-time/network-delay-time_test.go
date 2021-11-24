package _743_network_delay_time

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				4,
				2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
