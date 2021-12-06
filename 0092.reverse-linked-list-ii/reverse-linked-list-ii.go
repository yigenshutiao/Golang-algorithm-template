package _092_reverse_linked_list_ii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func reverseListNodes(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseListNodes(leftNode)

	pre.Next = rightNode
	leftNode.Next = cur

	return dummy.Next
}
