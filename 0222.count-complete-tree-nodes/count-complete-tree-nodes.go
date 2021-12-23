package _222_count_complete_tree_nodes

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func countNodes(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res += 1
		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)
	return res
}
