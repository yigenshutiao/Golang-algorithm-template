package _145_binary_tree_postorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func postorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}

	dfs(root)
	return res
}
