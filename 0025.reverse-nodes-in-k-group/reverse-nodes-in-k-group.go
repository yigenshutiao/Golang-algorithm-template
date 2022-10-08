package _025_reverse_nodes_in_k_group

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{Next: head}
	var pre, tail = dummy, head
	for tail != nil {

		for i := 0; i < k-1; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next //2)直接退出
			}
		} //1)完成p、h、t的定位

		tail.Next, tail = nil, tail.Next
		nh := reverse(head)
		pre.Next, head.Next, head, pre = nh, tail, tail, head
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode { //1)只用返回新头，因为输入的头就是新尾
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}
