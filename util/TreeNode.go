package util

// TreeNode is tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// IntSlice2TreeNode 利用 []int 生成 *TreeNode
func IntSlice2TreeNode(num []int) *TreeNode {
	n := len(num)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: num[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && num[i] != NULL {
			node.Left = &TreeNode{Val: num[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && num[i] != NULL {
			node.Right = &TreeNode{Val: num[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
