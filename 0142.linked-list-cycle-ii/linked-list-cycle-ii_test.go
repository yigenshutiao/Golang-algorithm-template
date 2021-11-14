package _142_linked_list_cycle_ii

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	pos  int
}{

	//{
	//	[]int{1},
	//	-1,
	//},
	//
	//{
	//	[]int{1, 2, 3},
	//	-1,
	//},

	{
		[]int{3, 2, 0, -4},
		1,
	},

	//{
	//	[]int{1, 2},
	//	0,
	//},

	// 可以有多个 testcase
}

func Test_detectCycle(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		head := util.Ints2ListWithCycle(tc.ints, tc.pos)
		var ans *ListNode
		if tc.pos >= 0 {
			ans = head.GetNodeWith(tc.ints[tc.pos])
		}
		ast.Equal(ans, detectCycle(head), "输入:%v", tc)
	}
}
