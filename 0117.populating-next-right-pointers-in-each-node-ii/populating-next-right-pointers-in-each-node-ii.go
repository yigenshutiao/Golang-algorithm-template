package _117_populating_next_right_pointers_in_each_node_ii

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

	queue := []*Node{root}
	nodes := [][]*Node{}
	for len(queue) > 0 {
		tmp := []*Node{}
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			tmp = append(tmp, node)
		}
		nodes = append(nodes, tmp)
	}

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[i])-1; j++ {
			nodes[i][j].Next = nodes[i][j+1]
		}
	}

	return root
}
