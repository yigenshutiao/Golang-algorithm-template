package _257_binary_tree_paths

import (
	"fmt"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
)

type TreeNode = util.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var traverse func(node *TreeNode, path string)
	traverse = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			v := path + fmt.Sprintf("%+v", node.Val)
			res = append(res, v)
			return
		}

		path = path + fmt.Sprintf("%+v", node.Val) + "->"
		if node.Left != nil {
			traverse(node.Left, path)
		}

		if node.Right != nil {
			traverse(node.Right, path)
		}
	}

	traverse(root, "")
	return res
}
