package _297_serialize_and_deserialize_binary_tree

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"strconv"
	"strings"
)

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
