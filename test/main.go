package test

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	left := lists[:len(lists)/2]
	right := lists[len(lists)/2:]

	return mergeTwoLists(mergeKLists(left), mergeKLists(right))
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy

	for list1 != nil && list2 != nil {
		if list2.Val > list1.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}

	if list2 != nil {
		dummy.Next = list2
	} else {
		dummy.Next = list1
	}
	return pre.Next
}

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var res []string

	var dfs func(left, right int, cur string, res *[]string)
	dfs = func(left, right int, cur string, res *[]string) {
		if left == 0 && right == 0 {
			*res = append(*res, cur)
			return
		}
		if left > right {
			return
		}

		if left > 0 {
			dfs(left-1, right, cur+"(", res)
		}

		if right > 0 {
			dfs(left, right-1, cur+")", res)
		}
	}

	dfs(n, n, "", &res)

	return res
}
