package offer07_zhong_jian_er_cha_shu_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	idx := getIndex(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func getIndex(order []int, target int) int {
	for i := 0; i < len(order); i++ {
		if order[i] == target {
			return i
		}
	}

	return -1
}
