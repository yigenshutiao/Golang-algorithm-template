package _530_minimum_absolute_difference_in_bst

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func getMinimumDifference(root *TreeNode) int {

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

	res := 99999999999

	for i := len(s) - 1; i > 0; i-- {
		t := s[i] - s[i-1]
		if t < res {
			res = t
		}
	}

	return res
}
