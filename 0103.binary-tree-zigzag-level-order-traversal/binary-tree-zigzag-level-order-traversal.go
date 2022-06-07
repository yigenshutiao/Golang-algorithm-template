package _103_binary_tree_zigzag_level_order_traversal

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}              // 结果数组，二维的
	curLevel := []*TreeNode{root} // 存放当前层的节点

	for len(curLevel) > 0 { // 直到当前层没有节点
		nextLevel := []*TreeNode{} // 存放下一层的节点
		curLevelVals := []int{}    // 存放当前层的节点值

		for _, node := range curLevel { // 遍历当前层的节点
			curLevelVals = append(curLevelVals, node.Val) // 节点值依次推入curLevelVals
			if node.Left != nil {                         // 如果有左子节点，推入nextLevel
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil { // 如果有右子节点，推入nextLevel
				nextLevel = append(nextLevel, node.Right)
			}
		}

		res = append(res, curLevelVals) // 遍历完当前层，curLevelVals推入res
		if len(res)%2 == 0 {            // 偶数层进行翻转
			reverse(curLevelVals)
		}
		curLevel = nextLevel // 下一轮遍历下一层节点
	}

	return res
}

func reverse(ints []int) { // 双指针法
	l, r := 0, len(ints)-1
	for l < r {
		ints[l], ints[r] = ints[r], ints[l]
		l++
		r--
	}
}
