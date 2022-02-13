package _560_subarray_sum_equals_k

// 前缀和版本
func subarraySum(nums []int, k int) int {
	res := 0
	pre := map[int]int{}
	pre[0] = 1

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, exist := pre[sum-k]; exist {
			res += pre[sum-k]
		}
		pre[sum]++
	}

	return res
}

// 暴力版本
func subarraySums(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}
