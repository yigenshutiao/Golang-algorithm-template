package _230_kth_smallest_element_in_a_bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

var tcs = []struct {
	pre, in []int
	k       int
	ans     int
}{

	//{
	//	[]int{1, 2},
	//	[]int{1, 2},
	//	2,
	//	2,
	//},
	//
	//{
	//	[]int{2, 1, 3},
	//	[]int{1, 2, 3},
	//	1,
	//	1,
	//},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		2,
		2,
	},

	//{
	//	[]int{2, 1, 3},
	//	[]int{1, 2, 3},
	//	3,
	//	3,
	//},

	// 可以有多个 testcase
}

func Test_kthSmallest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := util.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, kthSmallest(root, tc.k), "输入:%v", tc)
	}
}
