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
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
