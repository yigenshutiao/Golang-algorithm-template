package _104_maximum_depth_of_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		res++
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}

			queue = queue[1:]
		}
	}
	return res
}
