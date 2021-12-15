package _098_validate_binary_search_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func isValidBST(root *TreeNode) bool {
	min, max := -9999999999999, 9999999999999
	return traverse(root, min, max)
}

func traverse(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val < max && root.Val > min && traverse(root.Left, min, root.Val) && traverse(root.Right, root.Val, max)
}
