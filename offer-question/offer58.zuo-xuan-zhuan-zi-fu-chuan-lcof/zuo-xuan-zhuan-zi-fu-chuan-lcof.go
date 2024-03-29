package offer58_zuo_xuan_zhuan_zi_fu_chuan_lcof

func reverseLeftWords(s string, n int) string {
	// 先把前n个元素提取出来
	tmp := []byte{}
	b := []byte(s)
	for i := 0; i < n; i++ {
		tmp = append(tmp, b[i])
	}

	j := 0
	for i := n; i < len(b); i++ {
		b[j] = b[i]
		j++
	}

	x := 0
	for i := len(b) - n; i < len(b); i++ {
		b[i] = tmp[x]
		x++
	}
	return string(b)
}

func reverseLeftWords2(s string, n int) string {

	b := []byte(s)

	reverse(b[0:n])
	reverse(b[n:])
	reverse(b)

	return string(b)
}

func reverse(a []byte) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
