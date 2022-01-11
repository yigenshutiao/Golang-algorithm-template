package _199_binary_tree_right_side_view

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func rightSideView(root *TreeNode) []int {
	res := [][]int{}
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		res = append(res, tmp)
	}

	resp := []int{}

	for i := 0; i < len(res); i++ {
		resp = append(resp, res[i][len(res[i])-1])
	}
	return resp
}
