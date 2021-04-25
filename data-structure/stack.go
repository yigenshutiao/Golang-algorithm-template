package dataStructure

import "fmt"

var (
	tt  int
	stk = make(map[int]int)
)

func pop() {
	if tt > 0 {
		tt--
	}
}

func push(x int) {
	stk[tt] = x
	tt++
}

func boolEmpty() bool {
	if tt > 0 {
		return false
	}
	return true
}

func getStack() int {
	if boolEmpty() {
		return -1
	}
	return stk[tt]
}

func traverseStack() {
	for i := tt; i >= 0; i-- {
		fmt.Println(stk[tt])
	}
}

func main() {
	push(1)
	fmt.Println(tt, stk)
}
