package _114_flatten_binary_tree_to_linked_list

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// flatten 这个解法比较好理解...
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			// 若左孩子为空，不需要处理
			root = root.Right
		} else {
			// 当前节点同时具备俩子树
			pre := root.Left
			// 找到左孩子的最右边节点
			for pre.Right != nil {
				pre = pre.Right
			}
			// 把右孩子移到左边
			pre.Right = root.Right
			// 移好之后把左孩子放到右边
			root.Right = root.Left
			// 左节点置为空
			root.Left = nil
			// 处理下一个节点
			root = root.Right
		}
	}
}

// flatten 这种解法我估计第二次做也做不出来....
func flattenFk(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}
