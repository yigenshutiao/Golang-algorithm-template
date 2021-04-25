package double_pointer

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxSubSeq(num []int) int {
	//给定一个长度为n的整数序列,返回最长的不包含重复数字的连续子序列的长度
	res := 0
	count := make(map[int]int, len(num))
	j := 0
	for i := 0; i < len(num); i++ {
		count[num[i]]++
		for count[num[i]] > 1 {
			count[num[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	fmt.Print(maxSubSeq([]int{1, 2, 2, 3, 5}))
}
