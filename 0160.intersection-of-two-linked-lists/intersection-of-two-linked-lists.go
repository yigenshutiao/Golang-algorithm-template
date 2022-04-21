package _160_intersection_of_two_linked_lists

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		// 这里需要用p1 == nil 做判断，不能用p1.Next
		// 用 p1.Next 会重复
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func getIntersectionNodeIV(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != nil || curB != nil {
		if curA == curB {
			return curB
		}

		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return nil
}

// 二刷没记起来，记个搓点的打表法
func getIntersectionNodes(headA, headB *ListNode) *ListNode {
	visit := map[*ListNode]bool{}

	for node := headA; node != nil; node = node.Next {
		visit[node] = true
	}

	for node := headB; node != nil; node = node.Next {
		if visit[node] {
			return node
		}
	}
	return nil
}

// 三刷找到一个优美的解法
func getIntersectionNodeIII(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	lenA, lenB := 0, 0
	for nodeA != nil {
		nodeA = nodeA.Next
		lenA++
	}

	for nodeB != nil {
		nodeB = nodeB.Next
		lenB++
	}

	sub := 0
	slow, fast := &ListNode{}, &ListNode{}
	if lenA > lenB {
		sub = lenA - lenB
		slow, fast = headB, headA
	} else {
		sub = lenB - lenA
		slow, fast = headA, headB
	}

	for i := 0; i < sub; i++ {
		fast = fast.Next
	}

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
