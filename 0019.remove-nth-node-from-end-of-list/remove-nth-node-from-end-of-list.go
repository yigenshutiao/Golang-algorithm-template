package _019_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 链表题需要注意声明虚拟节点和类型问题...
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	phead := ListNode{-1, head}
	first := &phead

	for i := 0; i < n; i++ {
		first = first.Next
	}

	second := &phead
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return phead.Next
}
