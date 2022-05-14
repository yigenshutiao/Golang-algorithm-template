package _031_next_permutation

// nextPermutation, demo [1,2,3,4,6,5] => [1,2,3,5,4,6]
func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// 1. 找波谷：从右往左，寻找第一个比右邻居小的数，把它换到后面去
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 2. 找波谷尖：从右往左，寻找第一个比波谷大的数
	// 对于 [3，2，1]这样的情况，是没有波谷的，所以整体倒置
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		//将 A[i] 与 A[k] 交换
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 3. 把波谷后一位倒置，倒置后面序列 可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
	for x, y := j, len(nums)-1; x < y; {
		nums[x], nums[y] = nums[y], nums[x]
		x += 1
		y -= 1
	}

	return nums
}

func nextPermutationA(nums []int) {
	i, j := len(nums)-2, len(nums)-1

	// 从后往前找第一个降序：如果[i, j] 是降序，则继续，否则就是第一个升序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	k := len(nums) - 1
	// 这个条件说明找到了降序的序列
	if i >= 0 {
		// 从后向前，找第一个大于i的坐标，然后进行交换
		for nums[k] <= nums[i] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	// j后面的数字倒序
	for x, y := j, len(nums)-1; x < y; {
		nums[x], nums[y] = nums[y], nums[x]
		x++
		y--
	}
}
