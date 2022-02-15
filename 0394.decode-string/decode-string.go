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
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			result = str + strings.Repeat(result, count)
		} else {
			result += string(char)
		}
	}
	return result
}
