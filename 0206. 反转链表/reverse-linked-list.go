package _206_reverse_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	// 这里不能声明为 pre := &ListNode{}, 因为会初始化，这里只需要声明，不需要初始化
	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

//递归
func reverseListRecursion(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
