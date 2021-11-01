package main

import "fmt"

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	axis, pos := 0, len(a)-1
	for i := len(a) - 1; i >= 1; i-- {
		if a[i] > a[axis] {
			a[i], a[pos] = a[pos], a[i]
			pos--
		}
	}
	a[axis], a[pos] = a[pos], a[axis]

	quickSort(a[:pos])
	quickSort(a[pos+1:])
}

func main() {
	a := []int{5, 2, 3, 7, 6, 9, 0, 8, 4, 1}
	quickSort(a)
	fmt.Print(a)
}
