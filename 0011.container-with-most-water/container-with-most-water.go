package _011_container_with_most_water

// 暴力双重循环
func maxArea2(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res = max(res, (j-i)*min(height[i], height[j]))
		}
	}

	return res
}

// double pointer
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		res = max((r-l)*min(height[l], height[r]), res)
		// 移动短板，试图换取更长的高度
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
