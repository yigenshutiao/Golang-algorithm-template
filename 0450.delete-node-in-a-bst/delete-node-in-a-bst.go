package _450_delete_node_in_a_bst

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		tmp := root.Right

		for tmp.Left != nil {
			tmp = tmp.Left
		}

		tmp.Left = root.Left
		root = root.Right

		return root

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
