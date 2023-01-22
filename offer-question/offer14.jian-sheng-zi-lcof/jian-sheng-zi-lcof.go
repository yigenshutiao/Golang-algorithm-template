package offer14_jian_sheng_zi_lcof

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	x := n % 3
	a := n / 3
	if x == 0 {
		return int(math.Pow(3, float64(a)))
	} else if x == 1 {
		return int(4 * math.Pow(3, float64(a-1)))
	} else if x == 2 {
		return int(2 * math.Pow(3, float64(a)))
	}
	return -1
}
