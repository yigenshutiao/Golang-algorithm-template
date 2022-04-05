package _048_rotate_image

// rotate,解法比较特殊，先沿着135度对角线旋转，然后再reverse即可
func rotate(matrix [][]int) {
	// 沿着135度对角线旋转数组
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 然后对于每一行的元素进行reverse
	for _, nums := range matrix {
		reverse(nums)
	}
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
