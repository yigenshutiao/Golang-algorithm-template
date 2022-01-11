package _637_average_of_levels_in_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	nodes := [][]int{}
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

		nodes = append(nodes, tmp)
	}
	res := []float64{}
	for i := 0; i < len(nodes); i++ {
		t := 0
		for j := 0; j < len(nodes[i]); j++ {
			t += nodes[i][j]
		}
		res = append(res, float64(t)/float64(len(nodes[i])))
	}

	return res
}
