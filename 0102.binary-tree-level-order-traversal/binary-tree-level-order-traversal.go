package _102_binary_tree_level_order_traversal

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type TreeNode = util.TreeNode

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		res = append(res, []int{})
		for l := len(queue); l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			res[level] = append(res[level], node.Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
