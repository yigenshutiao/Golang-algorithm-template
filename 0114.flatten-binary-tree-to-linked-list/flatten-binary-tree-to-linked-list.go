package _114_flatten_binary_tree_to_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// flatten 这种解法我估计第二次做也做不出来....
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}
