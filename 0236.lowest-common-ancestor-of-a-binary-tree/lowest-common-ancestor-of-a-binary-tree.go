package _236_lowest_common_ancestor_of_a_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 搜到底了，返回
	if root == nil {
		return nil
	}
	// 搜到目标值了，返回当前节点
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 由于确认值的情况，所以只能用 != nil 来判断有相等的情况

	//当 left 和 right 同时不为空 ：说明 p, q 分列在 root的 异侧 （分别在 左 / 右子树），因此 root 为最近公共祖先，返回 root ；
	if left != nil && right != nil {
		return root
	}

	if left != nil && right == nil {
		return left
	}
	//当 left 为空 ，right 不为空 ：p,q 都不在 root 的左子树中，直接返回 right 。具体可分为两种情况：
	//p,q 其中一个在 root 的 右子树 中，此时 right 指向 p（假设为 p ）；
	//p,q 两节点都在 root 的 右子树 中，此时的 right 指向 最近公共祖先节点 ；
	if right != nil && left == nil {
		return right
	}

	// 当 left 和 right 同时为空 ：说明 root 的左 / 右子树中都不包含 p,q，返回 null
	if left == nil && right == nil {
		return nil
	}

	return nil
}
