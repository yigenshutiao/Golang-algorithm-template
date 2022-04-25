package _238_product_of_array_except_self

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) < 1 {
		return res
	}
	leftNums := make([]int, len(nums))
	leftNums[0] = 1
	for i := 1; i < len(nums); i++ {
		leftNums[i] = nums[i-1] * leftNums[i-1]
	}
	rigtNums := make([]int, len(nums))
	rigtNums[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rigtNums[i] = nums[i+1] * rigtNums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = leftNums[i] * rigtNums[i]
	}

	return res
}
