package _151_reverse_words_in_a_string

//1. 换成[]byte
//2. 反转整个
//3. 末尾加空格方便后续单词判断
//4. str[i] != ' ' && str[j] == ' '取单个单词并反转
//5. 拼接反转单词和空格
//6. 去除末尾多余空格
func reverseWords(s string) string {

	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}

	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}

	b = reverse(b)

	b = append(b, []byte(" ")...)
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] != ' ' && b[j] == ' ' {
				reverse(b[i:j])
				i = j
			}
		}
	}

	return string(b[:len(b)-1])
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
}
