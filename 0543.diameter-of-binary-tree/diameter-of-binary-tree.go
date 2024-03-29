package _543_diameter_of_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// diameterOfBinaryTree 直径就是二叉树最长的边
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := traverse(node.Left)
		right := traverse(node.Right)
		// 每个点都应该试着找最大值
		res = max(res, left+right)

		return max(left, right) + 1
	}

	traverse(root)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
