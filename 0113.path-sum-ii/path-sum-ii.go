package _113_path_sum_ii

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	tmp := []int{}

	var traverse func(node *TreeNode, count int, tmp []int)

	traverse = func(node *TreeNode, count int, tmp []int) {
		tmp = append(tmp, node.Val)
		count -= node.Val

		if node.Left == nil && node.Right == nil && count == 0 {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		if node.Left != nil {
			traverse(node.Left, count, tmp)
		}

		if node.Right != nil {
			traverse(node.Right, count, tmp)
		}
	}

	traverse(root, targetSum, tmp)

	return res
}
