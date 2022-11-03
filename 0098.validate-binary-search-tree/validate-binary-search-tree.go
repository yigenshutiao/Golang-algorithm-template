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
	// 中序遍历思路，pre是root的前一个节点
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftRes := dfs(node.Left)
		// 这里的Pre != nil 没想到
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		// 上面一卡，下面就可以赋值了
		pre = node

		rightRes := dfs(node.Right)

		return leftRes && rightRes

	}
	return dfs(root)
}
