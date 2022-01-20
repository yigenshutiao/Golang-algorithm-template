package _501_find_mode_in_binary_search_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_findMode(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     []int
	}{

		{
			[]int{1, 2, 1, 2},
			[]int{1, 1, 2, 2},
			[]int{1, 2},
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findMode(util.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
