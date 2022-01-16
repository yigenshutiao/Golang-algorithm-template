package _110_balanced_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	leftHigh := getHigh(root.Left)
	rightHigh := getHigh(root.Right)

	if abs(leftHigh-rightHigh) > 1 {
		return false
	}
	return true
}

func getHigh(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(getHigh(node.Left), getHigh(node.Right)) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
