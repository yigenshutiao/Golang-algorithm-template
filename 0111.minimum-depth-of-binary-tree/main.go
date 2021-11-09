package main

import (
	"fmt"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

func dfs(root *util.TreeNode, cur int, res *int) {
	if root == nil {
		return
	}

	cur += 1

	// 找到有效节点的终止条件
	if root.Left == nil && root.Right == nil {
		if cur < *res {
			*res = cur
		}
		return
	}

	dfs(root.Left, cur, res)
	dfs(root.Right, cur, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	res := 99999999
	dfs(root, 0, &res)

	return res
}

func minDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}

func main() {
	root := util.IntSlice2TreeNode([]int{3, 9, 20, 15, 7})
	fmt.Println(minDepth(root))
}
