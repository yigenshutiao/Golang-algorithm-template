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

func minDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	res := 99999999
	dfs(root, 0, &res)

	return res
}

func main() {
	root := util.IntSlice2TreeNode([]int{3, 9, 20, 15, 7})
	fmt.Println(minDepth(root))
}
