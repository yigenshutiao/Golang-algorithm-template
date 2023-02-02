package _001_two_sum

func twoSum(nums []int, target int) []int {
	for left := 0; left < len(nums); left++ {
		right := len(nums) - 1
		for right > left {
			if nums[right] == target-nums[left] {
				return []int{left, right}
			}
			right--
		}
	}
	return nil
}
