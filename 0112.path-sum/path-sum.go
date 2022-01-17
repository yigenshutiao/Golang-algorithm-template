package _112_path_sum

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if root.Left == nil && root.Right == nil && targetSum != 0 {
		return false
	}

	if root.Left != nil {
		if hasPathSum(root.Left, targetSum) {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, targetSum) {
			return true
		}
	}

	return false
}
