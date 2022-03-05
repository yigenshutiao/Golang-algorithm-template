package _461_hamming_distance

func hammingDistance(x int, y int) int {
	res := 0

	for x > 0 || y > 0 {
		res += (x & 1) ^ (y & 1)
		x = x >> 1
		y = y >> 1
	}

	return res
}
