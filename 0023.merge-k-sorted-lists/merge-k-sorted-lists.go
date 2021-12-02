package _023_merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	left := lists[0 : len(lists)/2]
	right := lists[len(lists)/2:]

	return mergeTwoLists(mergeKLists(left), mergeKLists(right))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 这里取指针类型的主要原因是后面代码有head = head.Next
	head := &ListNode{}
	pHead := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next // head 是已经被确定顺序的节点，head.Next 是后续操作要被确定顺序的节点
	}

	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return pHead.Next
}
