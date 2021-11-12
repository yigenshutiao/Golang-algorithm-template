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

	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}

	if l1.Val <= l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l1, l2.Next)
	}

	return head
}
