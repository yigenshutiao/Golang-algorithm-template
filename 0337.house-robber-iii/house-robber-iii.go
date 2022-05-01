package _337_house_robber_iii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func rob(root *TreeNode) int {
	// dp数组就两个元素，一个是0代表偷当前节点的最大值，1代表不偷当前节点的最大值
	// 当前节点依赖子节点，所以遍历顺序是后续遍历
	val := dfs(root)
	// 0代表选择当前节点，1代表不选择当前节点
	return max(val[0], val[1])
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	left, right := dfs(root.Left), dfs(root.Right)
	// 偷当前节点，左右孩子就不能偷
	choose := root.Val + left[1] + right[1]
	// 不偷当前节点，可以任意选择左右孩子的最大值
	notChoose := max(left[0], left[1]) + max(right[0], right[1])

	return []int{choose, notChoose}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
