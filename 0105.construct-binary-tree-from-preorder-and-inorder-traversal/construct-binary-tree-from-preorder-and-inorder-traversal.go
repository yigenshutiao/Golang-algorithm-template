package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// buildTree
// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	idx := findIndex(inorder, preorder[0])

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}

	return root
}

func findIndex(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}
