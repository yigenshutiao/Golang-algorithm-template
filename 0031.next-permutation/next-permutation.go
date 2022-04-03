package _031_next_permutation

func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// 从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是「小数」、「大数」
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		//将 A[i] 与 A[k] 交换
		nums[i], nums[k] = nums[k], nums[i]
	}

	//可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
	for x, y := j, len(nums)-1; x < y; {
		nums[x], nums[y] = nums[y], nums[x]
		x += 1
		y -= 1
	}

	return nums
}
