package _230_kth_smallest_element_in_a_bst

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

var i int
var res int

func kthSmallest(root *TreeNode, k int) int {
	i = 0
	res = -1
	run(root, k)
	return res
}

func run(root *TreeNode, k int) {
	if root == nil {
		return
	}
	run(root.Left, k)

	i++
	if i == k {
		res = root.Val
	}
	run(root.Right, k)
}
