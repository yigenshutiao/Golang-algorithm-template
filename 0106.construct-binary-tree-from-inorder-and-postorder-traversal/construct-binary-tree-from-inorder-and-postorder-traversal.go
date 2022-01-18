package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

//中序遍历 inorder = [9,3,15,20,7], 后序遍历 postorder = [9,15,7,20,3]
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	idx := findIndex(inorder, postorder[len(postorder)-1])

	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
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
