package _102_binary_tree_level_order_traversal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_Problem0102(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     [][]int
	}{

		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			[][]int{
				[]int{3},
				[]int{9, 20},
				[]int{15, 7},
			},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, levelOrder(util.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
