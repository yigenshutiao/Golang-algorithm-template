package main

import "fmt"

func prefixSum(a []int, l, r int) int {
	sum := []int{}
	sum = append(sum, 0) // 注意不能写成 sum[0]= 0
	for i := 1; i < len(a); i++ {
		sum = append(sum, sum[i-1]+a[i])
	}
	return sum[r] - sum[l-1]
}

func main() {
	a := []int{0, 2, 1, 3, 6, 4} // 下标从1开始, 0位置不参与运算
	fmt.Print(prefixSum(a, 2, 4))
}
