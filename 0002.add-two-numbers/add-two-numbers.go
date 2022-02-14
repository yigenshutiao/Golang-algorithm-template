package _002_add_two_numbers

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type ListNode = util.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10

		if head == nil {
			// 头结点只用一次
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 先记录尾节点的下一个节点，再蠕动尾节点
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
