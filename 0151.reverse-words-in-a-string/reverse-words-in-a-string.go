package _151_reverse_words_in_a_string

func reverseWords(s string) string {
	b := []byte(s)
	reverse(b)
	// 要加一个空格，不然最后会漏掉一个单词
	b = append(b, []byte{' '}...)
	res := []byte{}
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] != ' ' && b[j] == ' ' {
				t := reverse(b[i:j])
				res = append(res, t...)
				res = append(res, []byte{' '}...)
				i = j
			}
		}
	}
	res = res[:len(res)-1]
	return string(res)
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
}
