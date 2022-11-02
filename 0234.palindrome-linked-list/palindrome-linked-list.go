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

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := traverse(slow)

	newCur := head
	for right != nil {
		if newCur.Val != right.Val {
			return false
		}
		right = right.Next
		newCur = newCur.Next
	}

	return true

}

func traverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
