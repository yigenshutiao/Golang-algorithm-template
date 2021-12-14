package _038_binary_search_tree_to_greater_sum_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)

	return root
}
