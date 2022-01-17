package _112_path_sum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_Problem0112(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		sum     int
		ans     bool
	}{

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			22,
			true,
		},

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			26,
			true,
		},

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			10,
			false,
		},

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			100,
			false,
		},

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			11,
			false,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, hasPathSum(util.PreIn2Tree(tc.pre, tc.in), tc.sum), "输入:%v", tc)
	}
}
