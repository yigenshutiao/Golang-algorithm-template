package offer14_jian_sheng_zi_ii_lcof

func cuttingRope(n int) int {
	if n <= 3 {
		// 因为题目要求必须要剪>1段
		return n - 1
	}
	var b = n % 3
	var p = 1000000007
	var rem = 1
	// n = 3a+b, 对 x ^ (a-1) 循环求余
	for a := 1; a <= n/3-1; a++ {
		rem = (rem * 3) % p
	}
	if b == 0 {
		// 直接 补上最后一个3 即可
		return rem * 3 % p
	}
	if b == 1 {
		// 最后一个3+1 替换成 2 * 2
		return rem * 4 % p
	}
	// if b== 2, 3 * 2
	return rem * 6 % p
}
