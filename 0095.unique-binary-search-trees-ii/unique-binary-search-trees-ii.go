package _095_unique_binary_search_trees_ii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	res := traverse(1, n)
	return res
}

func traverse(min, max int) []*TreeNode {
	if min > max {
		// 不能写成[]*TreeNode{}
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}

	for i := min; i <= max; i++ {
		leftTrees := traverse(min, i-1)
		rightTrees := traverse(i+1, max)

		// 不能写成 for i:=0;i<len(leftTrees);i++ ...
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				node := &TreeNode{Val: i}
				node.Left = left
				node.Right = right
				res = append(res, node)
			}
		}
	}

	return res
}
