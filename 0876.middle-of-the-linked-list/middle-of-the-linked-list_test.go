package _876_middle_of_the_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

var tcs = []struct {
	head []int
	ans  []int
}{

	//{
	//	[]int{1, 2, 3, 4, 5},
	//	[]int{3, 4, 5},
	//},

	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{4, 5, 6},
	},

	// 可以有多个 testcase
}

func Test_middleNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := util.Ints2List(tc.head)
		actual := util.List2Ints(middleNode(head))
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}
