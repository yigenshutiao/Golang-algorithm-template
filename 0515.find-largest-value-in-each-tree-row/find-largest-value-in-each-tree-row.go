package _515_find_largest_value_in_each_tree_row

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func largestValues(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	nodes := [][]int{}
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
		nodes = append(nodes, tmp)
	}

	resp := []int{}
	for i := 0; i < len(nodes); i++ {
		if len(nodes[i]) > 1 {
			tmp := nodes[i][0]
			for j := 1; j < len(nodes[i]); j++ {
				tmp = max(tmp, nodes[i][j])
			}
			resp = append(resp, tmp)
		} else {
			resp = append(resp, nodes[i][0])
		}
	}

	return resp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
