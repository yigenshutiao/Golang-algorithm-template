package main

import "fmt"

func quickSort(q []int, l, r int) {
	if l >= r { // 终止条件 区间里面元素为1或没有时,返回
		return
	}
	var i, j = l-1, r+1 // 因为do while要先自增/自减
	var pivot = q[l]    // 这块在生产环境中会通过各种方式优化，比如在n个数中取中位数 // q[l]和j, j+1配套 q[r] 和 i-1, i配套。否则会出现各种边界问题
	for i < j {
		for { // do while 语法 在go中的写法
			i++ // 这步很关键，交换后指针要移动，避免没必要的交换
			if q[i] >= pivot {
				break
			}
		}
		for {
			j--
			if q[j] <= pivot {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i] // swap两个元素
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func main() {
	a := []int{5, 2, 3, 7, 6, 9, 0, 8, 4, 1}
	quickSort(a, 0, len(a)-1)
	fmt.Print(a)
}
