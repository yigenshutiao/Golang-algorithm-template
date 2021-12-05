package _206_reverse_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
