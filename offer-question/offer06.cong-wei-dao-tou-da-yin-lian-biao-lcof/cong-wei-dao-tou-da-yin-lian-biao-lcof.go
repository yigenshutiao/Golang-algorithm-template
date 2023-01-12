package offer06_cong_wei_dao_tou_da_yin_lian_biao_lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	res := []int{}

	cur := pre
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}
