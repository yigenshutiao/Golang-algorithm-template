package _019_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 声明一个虚拟节点，后面都从这个节点进行遍历即可
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
