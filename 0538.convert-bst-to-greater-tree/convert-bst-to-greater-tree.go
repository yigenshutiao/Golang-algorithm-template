package _538_convert_bst_to_greater_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	var num int
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Right)
		node.Val += num
		num = node.Val
		traverse(node.Left)
	}
	traverse(root)

	return root
}
