package _082_remove_duplicates_from_sorted_list_ii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if prev.Next == cur {
			prev = cur
		} else {
			prev.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		if pre.Next == head {
			pre = head
		} else {
			pre.Next = head.Next
		}
		head = head.Next
	}

	return dummy.Next
}
