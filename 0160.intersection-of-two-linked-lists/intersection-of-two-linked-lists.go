package _160_intersection_of_two_linked_lists

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		// 这里需要用p1 == nil 做判断，不能用p1.Next
		// 用 p1.Next 会重复
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
