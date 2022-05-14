package _206_reverse_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	// 这里不能声明为 pre := &ListNode{}, 因为会初始化，这里只需要声明，不需要初始化
	var pre *ListNode
	// 当前值要声明到head上...
	cur := head
	// head不会动，这里的变量是游标cur, 判断时应用cur进行判断
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//递归
func reverseListRecursion(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
