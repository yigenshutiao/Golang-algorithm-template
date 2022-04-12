package _102_binary_tree_level_order_traversal

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type TreeNode = util.TreeNode

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	// 这里需要用len判断，因为已经初始化了
	for len(queue) > 0 {
		l := len(queue)
		var tmp []int
		// 当前level的处理，queue处理当前层，存储下一层
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		res = append(res, tmp)
	}

	return res
}
