package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// buildTree
// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 先判断长度，返回空的情况
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 从preorder里面找到根节点的坐标，然后做拆分
	idx := getIdx(inorder, preorder[0])

	// 递归构造树
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
	return root
}

func getIdx(preorder []int, target int) int {
	for k, v := range preorder {
		if v == target {
			return k
		}
	}
	return -1
}
