package test

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"strconv"
	"strings"
)

func maximalSquare(matrix [][]byte) int {
	var corner int

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			t := min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1])
			if matrix[i][j] > '0' && t >= matrix[i][j] {
				matrix[i][j] = t + 1
			}
			a, _ := strconv.Atoi(string(matrix[i][j]))
			if a > corner {
				corner = a
			}
		}
	}

	return corner * corner
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

type TreeNode = util.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res string

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res += "null,"
			return
		}

		res += strconv.Itoa(node.Val) + ","
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var traverse func() *TreeNode

	traverse = func() *TreeNode {
		val := s[0]

		if val == "null" {
			s = s[1:]
			return nil
		}

		v, _ := strconv.Atoi(val)
		s = s[1:]

		return &TreeNode{
			Val:   v,
			Left:  traverse(),
			Right: traverse(),
		}
	}

	return traverse()
}
