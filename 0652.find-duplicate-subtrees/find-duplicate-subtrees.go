package _652_find_duplicate_subtrees

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"strconv"
)

type TreeNode = util.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[string]int{}
	res := []*TreeNode{}

	dfs(root, m, &res)

	return res
}

func dfs(root *TreeNode, m map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "*"
	}

	left := dfs(root.Left, m, res)
	right := dfs(root.Right, m, res)

	// 这里注意一下，必须用 strconv.Itoa
	key := strconv.Itoa(root.Val) + "_" + left + "_" + right
	m[key]++

	if m[key] == 2 {
		*res = append(*res, root)
	}

	return key
}
