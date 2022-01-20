package _501_find_mode_in_binary_search_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func findMode(root *TreeNode) []int {
	max := 1
	count := 1

	res := []int{}
	var pre *TreeNode

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)

		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}

		pre = node
		traverse(node.Right)
	}

	traverse(root)

	return res
}
