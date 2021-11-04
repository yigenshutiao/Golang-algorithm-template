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

	// 上面的代码，把数组分成独立的单元，每个单元个数为1
	// ---------- 以这里为分界线，把代码分成两块，上面的归并的拆分过程，下面是归并的合并过程 --------------
	// 下面是并的过程，把两个分开的序列，合成一个有序的序列

	// x是左序列的游标, y是右序列的游标, idx是merged数组的游标, merged 是本次合并后的res存放处
	x, y, idx, merged := 0, 0, 0, make([]int, len(arr))
	for x < len(left) && y < len(right) {
		if left[x] < right[y] {
			merged[idx] = left[x]
			x++
			idx++
		} else {
			merged[idx] = right[y]
			y++
			idx++
		}
	}
	// 把左序列剩下的元素放到merged中
	for x < len(left) {
		merged[idx] = left[x]
		idx++
		x++
	}
	// 把右序列剩下的元素放到merged中
	for y < len(right) {
		merged[idx] = right[y]
		idx++
		y++
	}
	copy(arr, merged)
}

func main() {
	//a := []int{5, 2, 3, 7, 6, 9, 0, 8, 4, 1}
	a := []int{4, 3, 2, 1}
	mergeSort(a)
	fmt.Print(a)
}
