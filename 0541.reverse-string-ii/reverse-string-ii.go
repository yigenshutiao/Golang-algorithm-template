package _541_reverse_string_ii

func reverseStr(s string, k int) string {
	res := []byte(s)

	for i := 0; i < len(res); i += (k * 2) {
		if len(res[i:]) >= k {
			reverse(res[i : i+k])
		} else {
			reverse(res[i:])
		}
	}
	return string(res)
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
