package _112_path_sum

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func traverse(root *TreeNode, target int) bool {
	if root.Left == nil && root.Right == nil && target == 0 {
		return true
	}

	if root.Left == nil && root.Right == nil && target != 0 {
		return false
	}

	if root.Left != nil {
		if hasPathSum(root.Left, target-root.Val) {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, target-root.Val) {
			return true
		}
	}

	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return traverse(root, targetSum)
}
