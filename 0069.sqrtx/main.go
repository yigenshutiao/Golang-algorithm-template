package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 0, x
	for r-l > 1 {
		mid := (l + r) >> 1
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}

func main() {

}
