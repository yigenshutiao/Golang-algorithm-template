package _581_shortest_unsorted_continuous_subarray

func findUnsortedSubarray(nums []int) int {
	left, right := -1, -1
	n := len(nums)
	minN, maxN := 999999999, -999999999
	for i := 0; i < n; i++ {
		// 从左往右走，应为升序(或者平)，若遇到降序，记录节点，取最后一个节点作为右边端点
		if nums[i] < maxN {
			right = i
		} else {
			maxN = nums[i]
		}
		// 从右往左走，应为降序(或者平)，若遇到升序，记录节点，取最后一个节点作为左边端点
		x := n - i - 1
		if nums[x] > minN {
			left = x
		} else {
			minN = nums[x]
		}
	}

	if right == -1 {
		return 0
	}

	return right - left + 1
}
