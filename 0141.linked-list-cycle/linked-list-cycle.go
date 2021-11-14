package _141_linked_list_cycle

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func hasCycle(head *ListNode) bool {

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
