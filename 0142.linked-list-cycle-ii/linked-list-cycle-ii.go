package _142_linked_list_cycle_ii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type ListNode = util.ListNode

//相遇时：
//slow指针走过的节点数为: x + y
//fast指针走过的节点数： x + y + n (y + z)，n为fast指针在环内走了n圈才遇到slow指针， （y+z）为 一圈内节点的个数
//因为fast指针是一步走两个节点，slow指针一步走一个节点， 所以 fast指针走过的节点数 = slow指针走过的节点数 * 2
//(x + y) * 2 = x + y + n (y + z)
//两边消掉一个（x+y）: x + y = n (y + z)
//因为我们要找环形的入口，那么要求的是x，因为x表示 头结点到 环形入口节点的的距离。
//所以我们要求x ，将x单独放在左面：x = n (y + z) - y
//当 n为1的时候，公式就化解为 x = z

// !!!这就意味着，从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点， 那么当这两个指针相遇的时候就是 环形入口的节点 !!!

// 记住相遇之后再走z距离即可
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {

			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}

	return nil
}
