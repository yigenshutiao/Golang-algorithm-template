package _429_n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {
			node := queue[0]
			for _, n := range node.Children {
				queue = append(queue, n)
			}
			queue = queue[1:]
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}

	return res
}
