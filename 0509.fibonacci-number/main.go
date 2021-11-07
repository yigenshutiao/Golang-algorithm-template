package main

import "fmt"

// 自底向顶进行计算，用的动态规划思想
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[len(arr)-1]
}

func main() {
	fmt.Println(fib(10))
}
