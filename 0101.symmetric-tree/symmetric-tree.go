package _101_symmetric_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 核心是要把left，right看成两棵树去后续遍历进行比较
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
