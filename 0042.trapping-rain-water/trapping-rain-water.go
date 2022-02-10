package _042_trapping_rain_water

// trap 按照列求每个雨水块的面积，优点是只需要计算高度即可
func trap(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		if i == 0 || i == len(height)-1 {
			continue
		}

		leftHigh, rightHigh := height[i], height[i]

		for l := i; l >= 0; l-- {
			if height[l] > leftHigh {
				leftHigh = height[l]
			}
		}

		for r := i + 1; r <= len(height)-1; r++ {
			if height[r] > rightHigh {
				rightHigh = height[r]
			}
		}

		s := min(leftHigh, rightHigh) - height[i]

		res += s
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
