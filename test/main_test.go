package test

import (
	"fmt"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {

	tcs := struct {
		pre, in []int
	}{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 1, 4, 3, 5},
	}

	ser := Constructor()

	deser := Constructor()

	data := ser.serialize(util.PreIn2Tree(tcs.pre, tcs.in))

	ans := deser.deserialize(data)

	fmt.Println(ans)
}
