package _215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间 [left, right] 执行 partition 操作，返回 nums[left]（选取的第一个作为pivot） 排序以后应该在的位置
	partition := func(nums []int, left, right int) int {
		pivot := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			// 小于 pivot 的元素都被交换到前面
			if nums[i] < pivot {
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		// 在之前遍历的过程中，满足 [left+1, j] < pivot，并且 (j, i] >= pivot
		nums[j], nums[left] = nums[left], nums[j]
		// 交换以后 [left, j-1] < pivot, nums[j] = pivot, [j+1, right] >= pivot
		return j
	}

	n := len(nums)
	// 转换一下，第 k 大元素的索引是 n - k
	left, right, target := 0, n-1, n-k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}
