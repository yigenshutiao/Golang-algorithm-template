package _559_maximum_depth_of_n_ary_tree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	res := 0
	queue := []*Node{root}

	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[0]

			for i := 0; i < len(node.Children); i++ {
				queue = append(queue, node.Children[i])
			}
			queue = queue[1:]
		}
		res++
	}

	return res
}
