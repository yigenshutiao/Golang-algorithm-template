package bigNumOperation

import (
	"fmt"
)

func div(a []int, b int) []int {
	if b == 0 {
		return nil
	}
	res := []int{}
	r := 0
	for i := len(a) - 1; i >= 0; i-- {
		r = r*10 + a[i]
		res = append(res, r/b)
		r %= b
	}
	// reverse arr
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	// 去掉答案中的前导0
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func callDiv() {

	a := "120"
	b := 0

	A := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i])-'0')
	}

	c := div(A, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}

}
