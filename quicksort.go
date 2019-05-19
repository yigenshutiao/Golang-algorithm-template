package Go_Algorithm_Template

func quickSort(q []int, l, r int) {
	if l >= r { // 终止条件
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
