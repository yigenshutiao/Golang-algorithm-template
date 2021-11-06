package main

import "fmt"

//threeSqrt 给定一个浮点数 n，求它的三次方根
func threeSqrt(x float64) float64 {
	l, r := 0.0, x
	for (r - l) > 1e-6 {
		mid := (l + r) / 2
		if (mid * mid * mid) > x {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	fmt.Println(threeSqrt(1000.00))
}
