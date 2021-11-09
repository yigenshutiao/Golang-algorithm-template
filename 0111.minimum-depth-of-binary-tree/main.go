package main

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode, cur int, res *int) {
	if root == nil {
		return
	}

	cur += 1

	if root.Left == nil && root.Right == nil {
		if cur < *res {
			*res = cur
		}
		return
	}

	bfs(root.Left, cur, res)
	bfs(root.Right, cur, res)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 99999999
	bfs(root, 0, &res)

	return res
}

func main() {
	util.IntSlice2TreeNode()
}
