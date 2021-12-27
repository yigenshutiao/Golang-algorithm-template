package _094_binary_tree_inorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
