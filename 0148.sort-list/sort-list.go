package _148_sort_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	// 这里要做初始化, 不然后面取Next指针可能会nil异常
	pre := new(ListNode)

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 不要忘记截断next指针
	pre.Next = nil
	list1 := sortList(head)
	list2 := sortList(slow)

	return merge(list1, list2)
}

func merge(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}
