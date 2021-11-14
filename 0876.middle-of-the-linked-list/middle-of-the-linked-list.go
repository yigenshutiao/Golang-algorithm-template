package _876_middle_of_the_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = util.ListNode

func middleNode(head *ListNode) *ListNode {

	one, second := head, head

	// second != nil 为了判断快指针走到末尾的情况
	for second != nil && second.Next != nil {
		one = one.Next
		second = second.Next.Next
	}

	return one
}
