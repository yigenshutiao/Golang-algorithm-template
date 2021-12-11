package _116_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)

	return root
}

func connectTwoNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)

	// 这里要注意一下
	connectTwoNode(node1.Right, node2.Left)
}
