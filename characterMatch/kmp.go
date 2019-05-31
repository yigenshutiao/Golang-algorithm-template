package characterMatch

import "fmt"

func kmp(s, pattern string) int {
	n, m := len(s), len(pattern)
	if n < m {
		return -1
	}

	next := make([]int, m)
	// next数组中的值初始化为-1
	for index := range next {
		next[index] = -1
	}
	//求next数组中的值
	for i := 1; i < m-1; i++ { // i从1开始,因为第一个字符如果比较失败了,需重新开始匹配 // i取不到m-1的值, 因为取到m-1整个字符串就相等了
		j := next[i-1] // 前i-1的值是之前循环中比较过的,这里j初始化为next[i-1]

		for pattern[j+1] != pattern[i] && j != -1 { // 因为这里是pattern[i]和pattern[j+1]进行比较
			j = next[j] // 所以这里j是退回到next[j]的位置再进行循环比较
		}

		if pattern[j+1] == pattern[i] { //因为每次循环只会新增一个字符,所以这里用if判断一个新字母即可.
			j++ // 如果相等则j++
		}

		next[i] = j // 当前的取值
	}
	// 匹配的过程
	j := 0 //模式串从0下标开始匹配
	for i := 0; i < n; i++ {
		for j > 0 && s[i] != pattern[j] { // j>0意为j没有退回起点 //s[i] != pattern[j]意为两个字符出现不匹配的情况
			j = next[j-1] + 1 // pattern[j]和s[i]不一致,说明前next[j-1]是匹配的,所以移动next[j-1]位;因为s[i]要继续和pattern[j]进行比较,所以j还需加1
		}

		if s[i] == pattern[j] {
			if j == m-1 { //因为j从下标0开始,所以m需减1,两者相等说明循环了len(m)次
				return i - m + 1
			}
			j++ //否则继续判断下一个字符
		}
	}
	return -1
}

func main() {
	a := "ababaababacd"
	b := "ababac"
	fmt.Print(kmp(a, b))
}
