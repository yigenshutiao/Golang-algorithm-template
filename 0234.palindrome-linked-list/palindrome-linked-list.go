package _234_palindrome_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = util.ListNode

func traverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	// 反转链表并赋给新节点
	right := traverse(slow)

	cur := head
	for right != nil {
		if cur.Val != right.Val {
			return false
		}
		right = right.Next
		cur = cur.Next
	}

	return true
}
