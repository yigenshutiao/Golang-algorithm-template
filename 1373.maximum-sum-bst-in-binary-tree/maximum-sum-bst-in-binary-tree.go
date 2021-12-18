package _373_maximum_sum_bst_in_binary_tree

import "github.com/yigenshutiao/Golang-algorithm-template/util"

type TreeNode = util.TreeNode

var res int

func maxSumBST(root *TreeNode) int {
	res = 0
	traverse(root)
	return res
}

func traverse(root *TreeNode) (isValid bool, minVal, maxVal, sum int) {
	if root == nil {
		return true, 99999999999, -99999999999, 0
	}

	leftValid, leftMin, leftMax, leftSum := traverse(root.Left)
	rightValid, rightMin, rightMax, rightSum := traverse(root.Right)

	if leftValid && rightValid && root.Val > leftMax && root.Val < rightMin {
		isValid = true
		minVal = getMin(root.Val, leftMin)
		maxVal = getMax(root.Val, rightMax)
		sum = root.Val + leftSum + rightSum
		if sum > res {
			res = sum
		}
	} else {
		isValid = false
	}
	return
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
