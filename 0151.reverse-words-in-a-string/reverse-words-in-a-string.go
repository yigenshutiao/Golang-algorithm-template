package _151_reverse_words_in_a_string

//1.换成[]byte
//2.反转整个
//3.末尾加空格方便后续单词判断
//4.str[i]!=' '&&str[j]==' '取单个单词并反转
//5.拼接反转单词和空格
//6.去除末尾多余空格
func reverseWords(s string) string {
	b := []byte(s)
	b = reverse(b)

	b = append(b, []byte(" ")...)
	var res []byte
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] != ' ' && b[j] == ' ' {
				t := reverse(b[i:j])
				res = append(res, t...)
				res = append(res, []byte(" ")...)
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
