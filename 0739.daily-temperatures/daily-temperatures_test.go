package _739_daily_temperatures

import (
	"reflect"
	"testing"
)

func Test_dailyTemperature(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				[]int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperature(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
