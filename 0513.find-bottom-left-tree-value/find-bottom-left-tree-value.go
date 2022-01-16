package _513_find_bottom_left_tree_value

import (
	"fmt"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type TreeNode = util.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	res := [][]int{}

	for len(queue) > 0 {
		t := []int{}
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			t = append(t, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res = append(res, t)
	}

	fmt.Println(res)

	return res[len(res)-1][0]
}
