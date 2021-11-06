## 二分问题

二分问题指的是:

> 把一段区间一分为二，一半区间满足某种性质，另一半不满足。此时可以用二分法寻找边界。
> 二分本质上是在找边界

思想:

> 在一个区间内部二分答案，每次选答案所在的区间进行处理。当区间为1的时候，就是答案。

二分模板分为整数模板和浮点数模板,其中整数模板有两种

整数二分模板：
```go
// 检查x是否满足某种性质
func check(x int) bool

// 区间[l, r]被划分成[l, mid]和[mid + 1, r]时使用：
func bsearch(l, r int) int {
	for l < r {
		mid := (l+r) >> 1
		if check(mid) {  // check()判断mid是否满足性质
		    r = mid	
        } else {
            l = mid + 1	
        }
	}
	return l
}

// 区间[l, r]被划分成[l, mid - 1]和[mid, r]时使用：
func bsearch(l, r int) int {
    for l < r {
        mid := (l+r+1) >> 1
        if check(mid) {  // check()判断mid是否满足性质
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l
}
```

```go
// 浮点数二分模板
// check 检查x是否满足某种性质
func check(x float64) bool

func bearchSearch(l,r float64) float64 {
	eps := 1e-6 // eps 表示精度，取决于题目对精度的要求
	for (r - l) > eps {
	    mid = (l+r)/2
		if 	check(mid) {
		    r = mid	
        } else {
        	l = mid
        }
    }
    return l
}
```

相关LeetCode题:
```
 69. 求x的平方根 
 162. 寻找峰值
```

