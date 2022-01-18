package _654_maximum_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

// [3,2,1,6,0,5]
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	idx := getMax(nums)

	root := &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[0:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1 : len(nums)]),
	}

	return root
}

func getMax(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}

	return res
}
