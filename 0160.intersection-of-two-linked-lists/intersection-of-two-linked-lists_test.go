package _160_intersection_of_two_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	a, b   []int
	ea, eb int
}{

	{
		[]int{4, 1, 8, 4, 5},
		[]int{5, 0, 1, 8, 4, 5},
		2,
		3,
	},

	//{
	//	[]int{0, 9, 1, 2, 4},
	//	[]int{3, 2, 4},
	//	3,
	//	1,
	//},
	//
	//{
	//	[]int{2, 6, 4},
	//	[]int{1, 5},
	//	3,
	//	2,
	//},

	// 可以有多个 testcase
}

// head must Not be nil
func tailOf(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func Test_getIntersectionNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		tail := util.Ints2List(tc.a[tc.ea:])

		headA := util.Ints2List(tc.a[:tc.ea])
		tailA := tailOf(headA)
		tailA.Next = tail

		headB := util.Ints2List(tc.b[:tc.eb])
		tailB := tailOf(headB)
		tailB.Next = tail

		ast.Equal(tail, getIntersectionNode(headA, headB), "输入:%v", tc)
	}
}
