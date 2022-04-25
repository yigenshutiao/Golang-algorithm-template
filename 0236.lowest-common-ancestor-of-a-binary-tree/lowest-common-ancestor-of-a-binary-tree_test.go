package _236_lowest_common_ancestor_of_a_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

var root = util.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, util.NULL, util.NULL, 7, 4})
var tcs = []struct {
	p, q, ans *TreeNode
}{

	//{
	//	util.GetTargetNode(root, 5),
	//	util.GetTargetNode(root, 1),
	//	util.GetTargetNode(root, 3),
	//},
	//
	//{
	//	util.GetTargetNode(root, 4),
	//	util.GetTargetNode(root, 7),
	//	util.GetTargetNode(root, 2),
	//},

	{
		util.GetTargetNode(root, 4),
		util.GetTargetNode(root, 5),
		util.GetTargetNode(root, 5),
	},

	//{
	//	util.GetTargetNode(root, 5),
	//	util.GetTargetNode(root, 4),
	//	util.GetTargetNode(root, 5),
	//},
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		node := lowestCommonAncestor(root, tc.p, tc.q)
		ast.Equal(tc.ans, node, "p=%d,q=%d", tc.p.Val, tc.q.Val)
	}
}
