package dataStructure

import "fmt"

var (
	index int
	val    = make(map[int]int)
	l        = make(map[int]int) //存左指针
	r        = make(map[int]int) //存右指针
)

func init() {
	r[0] = 1
	l[1] = 0
	index = 2

}

// 在节点k的右边插入一个值x
// 若想在k的左边插入节点, 则insert(l[k], x)
func insert(k, x int) {
	val[index] = x
	l[index] = k
	r[index] = r[k]
	l[r[k]] = index
	r[k] = index
	index++
}

// 删除节点k
func removeNode(k int) {
	r[l[k]] = r[k]
	l[r[k]] = l[k]
}

func traverses() {
	for i := r[0]; i != 1; i = r[i] {
		fmt.Println(val[i])
	}
}
