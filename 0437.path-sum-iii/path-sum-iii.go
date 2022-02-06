package _437_path_sum_iii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func pathSum(root *TreeNode, targetSum int) int {

	res := 0
	if root == nil {
		return res
	}

	var traverse func(node *TreeNode, cur int)
	traverse = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}

		if cur == node.Val {
			res++
		}

		traverse(node.Left, cur-node.Val)
		traverse(node.Right, cur-node.Val)
		return
	}

	traverse(root, targetSum)

	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
