package _235_lowest_common_ancestor_of_a_binary_search_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
