package offer29_shun_shi_zhen_da_yin_ju_zhen_lcof

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1 //初始边界

	// 循环条件，下大于等于上，右大于等于左
	for bottom >= top && right >= left {
		// top行的从左到右遍历
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		// top行遍历完了，top往下移动
		top++

		// right列的从上到下遍历
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		// right列遍历完了，right向左移动
		right--

		// 这里一定要做一个判断，因为如果只剩下一个元素就会出错
		if left > right || top > bottom {
			break
		}

		// bottom行的从右往左遍历
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}

		//bottom往上移动
		bottom--

		for i := bottom; i >= top; i-- { //left列的从下到上遍历
			res = append(res, matrix[i][left])
		}
		//left向右移动
		left++
	}
	return res
}
