package _101_symmetric_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func isSymmetric(root *TreeNode) bool {
	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		// 先判断空
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			return false
		}
		// 看当前值是否相等
		if node1.Val != node2.Val {
			return false
		}

		// 看下一层是否相等
		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}

	return dfs(root.Left, root.Right)
}
