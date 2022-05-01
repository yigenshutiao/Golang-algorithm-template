package _394_decode_string

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			strStack = append(strStack, result)
			result = ""

			numStack = append(numStack, num)
			num = 0
		} else if char == ']' {
			// 这里的数字和字母都要重新声明，不能影响前面的结果
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			// 最终结果 = pope出来的字段 + 当前的字符 * 重复次数，这个公式要注意一下
			result = str + strings.Repeat(result, count)
		} else {
			result += string(char)
		}
	}
	return result
}
