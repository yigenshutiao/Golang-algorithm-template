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

	left := lists[:len(lists)/2]
	right := lists[len(lists)/2:]

	return mergeTwoList(mergeKLists(left), mergeKLists(right))
}

func mergeTwoList(head1, head2 *ListNode) *ListNode {
	if head1 == nil && head2 == nil {
		return nil
	}

	dummy := new(ListNode)
	pre := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			dummy.Next = head1
			head1 = head1.Next
		} else {
			dummy.Next = head2
			head2 = head2.Next
		}
		dummy = dummy.Next
	}

	if head1 != nil {
		dummy.Next = head1
	} else {
		dummy.Next = head2
	}

	return pre.Next
}
