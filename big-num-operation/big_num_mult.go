package big_num_operation

import "fmt"

func mult(a []int, b int) []int {
	res := []int{}
	t := 0
	for i := 0; i < len(a) || t > 0; i++ {
		if i < len(a) {
			t += a[i] * b
		}
		res = append(res, t%10)
		t /= 10
	}
	return res
}

func callMult() {
	a := "9999"
	b := 9

	A := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i])-'0')
	}
	c := mult(A, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
}
