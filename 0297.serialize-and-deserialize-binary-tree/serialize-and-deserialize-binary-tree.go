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
func (Codec) serialize(root *TreeNode) string {
	res := &strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			res.WriteString("null,")
			return
		}

		res.WriteString(strconv.Itoa(root.Val) + ",")
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return res.String()
}

// Deserializes your encoded data to tree.
func (Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var traverse func() *TreeNode
	traverse = func() *TreeNode {
		if s[0] == "null" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		s = s[1:]
		return &TreeNode{val, traverse(), traverse()}
	}

	return traverse()
}
