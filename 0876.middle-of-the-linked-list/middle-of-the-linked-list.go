package _876_middle_of_the_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func middleNode(head *ListNode) *ListNode {

	one, second := head, head

	// second != nil 为了判断快指针走到末尾的情况
	// 当second 已经走到末尾的nil, 若再判断second的Next，就会panic
	for second != nil && second.Next != nil {
		one = one.Next
		second = second.Next.Next
	}

	return one
}
