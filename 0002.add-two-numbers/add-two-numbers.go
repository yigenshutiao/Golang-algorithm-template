package _002_add_two_numbers

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type ListNode = util.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 记住要起一个新链表即可
	dummy := new(ListNode)
	pre := dummy

	var s, v int
	// 两个节点有一个不为空即可，每次拿到当前值和是否进位
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + s
		s, v = sum/10, sum%10

		cur := &ListNode{Val: v}
		dummy.Next = cur
		dummy = dummy.Next
	}

	if s > 0 {
		cur := &ListNode{Val: s}
		dummy.Next = cur
	}

	return pre.Next
}
