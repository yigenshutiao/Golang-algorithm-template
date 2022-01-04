package _019_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	slow, fast := pre, pre
	// 一共5个节点，n = 3，n+1=4，再走1步就能找到n的前置节点，这是走n+1步的原因
	// 1. 快指针走n+1步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	// 2. 快慢指针同时走，让慢指针走到n节点的pre
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 3. 开始删除
	slow.Next = slow.Next.Next

	return pre.Next
}
