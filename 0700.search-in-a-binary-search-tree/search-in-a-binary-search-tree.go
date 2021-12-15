package _700_search_in_a_binary_search_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		root = searchBST(root.Left, val)
	} else if root.Val < val {
		root = searchBST(root.Right, val)
	}

	return root
}
