package doublePointer

import "fmt"

func getChar(str string) {
	// 打印出str中的每个子字符串, 并用","分隔, str以字母开头, 子字符串之间只有一个空格
	n := len(str)
	for i := 0; i < n; i++ {
		j := i
		for j < n && str[j] != ' ' {
			j++
		}
		for k := i; k < j; k++ {
			fmt.Print(string(str[k]))
		}
		fmt.Print(",")
		i = j
	}

}

func main() {
	getChar("adb adha ahdjka")
}
