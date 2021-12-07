package _025_reverse_nodes_in_k_group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	k    int
	ans  []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		3,
		[]int{3, 2, 1, 4, 5},
	},

	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{2, 1, 4, 3, 5},
	},

	// 可以有多个 testcase
}

func Test_reverseKGroup(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := util.Ints2List(tc.head)
		ans := util.Ints2List(tc.ans)
		ast.Equal(ans, reverseKGroup(head, tc.k), "输入:%v", tc)
	}
}
