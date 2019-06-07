package main

import "fmt"

var (
	head, idx int                 // head 头结点下标  idx 当前用到的点
	e         = make(map[int]int) // val[i] 表示节点i的值
	ne        = make(map[int]int) // ne[i] 表示节点i的next指针是多少
)

//初始化
func inits() {
	head = -1
	idx = 0
}

// 将x插到头结点
func addToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// 将x插到下标为k的点的后面
func add(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

// 将下标是k的点后面的点删掉
func remove(k int) {
	ne[k] = ne[ne[k]]
}

// 遍历链表
func traverse(head int) []int {
	res := []int{}
	for i := head; i != -1; i = ne[i] {
		res = append(res, e[i])
	}
	return res
}

func main() {
	inits()
	fmt.Println("head:", head, "idx:", idx, "e:", e, "ne:", ne)

	addToHead(5)
	fmt.Println("head:", head, "idx:", idx, "e:", e, "ne:", ne)

	add(0, 6)
	fmt.Println("head:", head, "idx:", idx, "e:", e, "ne:", ne)

	add(1, 4)
	fmt.Println("head:", head, "idx:", idx, "e:", e, "ne:", ne)

	remove(0)
	fmt.Println("head:", head, "idx:", idx, "e:", e, "ne:", ne)

	linklist := traverse(head)
	fmt.Println(linklist)
}
