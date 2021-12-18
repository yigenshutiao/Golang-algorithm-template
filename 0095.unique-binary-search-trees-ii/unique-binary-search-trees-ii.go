package _095_unique_binary_search_trees_ii

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type TreeNode = util.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	res := traverse(1, n)
	return res
}

func traverse(min, max int) []*TreeNode {
	// min = max 还有一个节点，大于才需要返回
	if min > max {
		// 不能写成 []*TreeNode{}, 这个相当于 var s []*TreeNode, 只有声明，没有赋值
		// 如果不赋值，后面就不能遍历了
		return []*TreeNode{nil}
	}
	// 每次递归的 res 都是不同的
	res := []*TreeNode{}

	// 这里要去<=的情况，因为有两个节点相等的情况
	// i 是当前的节点， [min, i-1]是左节点，[i+1,max]是右节点
	for i := min; i <= max; i++ {
		leftTrees := traverse(min, i-1)
		rightTrees := traverse(i+1, max)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				node := &TreeNode{Val: i}
				node.Left = left
				node.Right = right
				res = append(res, node)
			}
		}
	}
	// 返回当前情况能组成的所有树
	return res
}
