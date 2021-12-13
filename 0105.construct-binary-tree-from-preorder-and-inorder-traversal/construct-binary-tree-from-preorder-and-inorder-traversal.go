package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// buildTree
// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := findIndex(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return root
}

func findIndex(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
