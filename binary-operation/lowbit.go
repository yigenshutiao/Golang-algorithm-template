package binary_operation

import "fmt"

func lowbit(x int) int {
	return x & -x
}

func getCount(x int) int {
	//给定一个长度为n的数列,请你求出数列中每个数的二进制表示中1的个数
	res := 0
	for x > 0 {
		x -= lowbit(x)
		res++
	}
	return res
}

func main() {
	fmt.Print(getCount(15))
}
