package main

import (
	"fmt"
)

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	left, right := arr[:len(arr)/2], arr[len(arr)/2:]

	mergeSort(left)
	mergeSort(right)

	// merge
	x, y, idx, merged := 0, 0, 0, make([]int, len(arr))
	for x < len(left) && y < len(right) {
		if left[x] < right[y] {
			merged[idx] = left[x]
			x++
			idx++
			continue
		}
		if left[x] >= right[y] {
			merged[idx] = right[y]
			y++
			idx++
			continue
		}
	}
	for x < len(left) {
		merged[idx] = left[x]
		idx++
		x++
	}
	for y < len(right) {
		merged[idx] = right[y]
		idx++
		y++
	}
	copy(arr, merged)
}

func main() {
	a := []int{5, 2, 3, 7, 6, 9, 0, 8, 4, 1}
	mergeSort(a)
	fmt.Print(a)
}
