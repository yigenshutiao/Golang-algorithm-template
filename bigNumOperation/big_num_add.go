package bigNumOperation

import "fmt"

func bigNumAdd(a, b []int) []int {
	res := []int{}
	t := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		res = append(res, t%10)
		t /= 10
	}
	if t == 1 {
		res = append(res, 1)
	}
	return res
}

func callAdd() {
	a := "89"
	b := "31"
	arrA, arrB := []int{}, []int{}
	for i := len(a) - 1; i >= 0; i-- {
		arrA = append(arrA, int(a[i])-'0')
	}
	for i := len(b) - 1; i >= 0; i-- {
		arrB = append(arrB, int(b[i])-'0')
	}
	c := bigNumAdd(arrA, arrB)

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
}
