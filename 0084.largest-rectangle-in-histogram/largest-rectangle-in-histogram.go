package _084_largest_rectangle_in_histogram

// largestRectangleAreas 此方法超时了...
func largestRectangleAreas(heights []int) int {
	var res int

	for i := 0; i < len(heights); i++ {
		leftH, rightH := i, i
		for m := i - 1; m >= 0; m-- {
			if heights[m] >= heights[i] {
				leftH = m
			} else {
				break
			}
		}

		for n := i + 1; n < len(heights); n++ {
			if heights[n] >= heights[i] {
				rightH = n
			} else {
				break
			}
		}
		width := rightH - leftH + 1
		sum := width * heights[i]
		res = max(res, sum)
	}

	return res
}

// heights = [2,1,5,6,2,3] dp方法
func largestRectangleArea(heights []int) int {
	var res int

	// leftIdx 当前坐标左边第一个小于其高度的坐标， rightIdx 当前坐标右边第一个小于其的坐标
	leftIdx, rightIdx := make([]int, len(heights)), make([]int, len(heights))

	leftIdx[0] = -1
	for i := 1; i < len(heights); i++ {
		// t：最终存第一个小于i的坐标，先预设其小于左边第一个坐标要小于 hei[i]
		t := i - 1
		for t >= 0 && heights[t] >= heights[i] {
			// 这里是在移动t
			t = leftIdx[t]
		}
		leftIdx[i] = t
	}

	rightIdx[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[t] >= heights[i] {
			t = rightIdx[t]
		}
		rightIdx[i] = t
	}

	for i := 0; i < len(heights); i++ {
		sum := heights[i] * (rightIdx[i] - leftIdx[i] - 1)
		res = max(res, sum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
