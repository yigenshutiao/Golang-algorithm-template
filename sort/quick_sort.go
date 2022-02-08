package main

import "fmt"

// quickSort 采用了分治的策略，将原问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题。
// 快排的基本思想是：
// 在序列中找一个划分值，通过一趟排序将未排序的序列排序成 独立的两个部分，
// 其中左边部分序列都比划分值小，右边部分的序列比划分值大，此时划分值的位置已确认，
// 然后再对这两个序列按照同样的方法进行排序，从而达到整个序列都有序的目的。
func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	// axis是坐标值，pos是其最终应处位置
	axis, pos := 0, len(a)-1
	for i := len(a) - 1; i >= 1; i-- {
		if a[i] > a[axis] {
			// a[i]是大于坐标的值，a[pos]是小于坐标的值
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
