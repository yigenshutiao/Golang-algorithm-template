package _098_validate_binary_search_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// 中序遍历BST应该得到一个递增序列
func isValidBSTTraverse(root *TreeNode) bool {
	s := []int{}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		s = append(s, node.Val)
		traverse(node.Right)
	}

	traverse(root)

	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			return false
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode) bool
	var pre *TreeNode
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftRes := dfs(node.Left)
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node

		rightRes := dfs(node.Right)

		return leftRes && rightRes

	}
	return dfs(root)
}
