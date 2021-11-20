## 二分问题

二分问题指的是:

> 把一段区间一分为二，一半区间满足某种性质，另一半不满足。此时可以用二分法寻找边界。
> 二分本质上是在找边界

思想:

> 在一个区间内部二分答案，每次选答案所在的区间进行处理。当区间为1的时候，就是答案。

二分模板分为整数模板和浮点数模板,其中整数模板有两种

整数二分模板：
```go
// 寻找左边界的二分查找
// 返回l，代表在nums中，有l个元素小于target

func bsearch(nums []int,target int) int {
    l, r := 0, len(num)-1
    // 终止条件 l == r
    for l < r {
        mid := (l + r) >> 1
        if num[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

// 寻找右边界的二分查找
func bsearch(num []int, target int) int {
    l, r := 0, len(num)-1
    for l < r {
    // 当 l = r - 1  时，mid 若不加1，则等于l，区间为[l, r], 没有变化, 造成死循环
    // 只有 +1,才能保证mid = r，跳出终止条件
        mid := (l + r + 1) >> 1
        if num[mid] <= target {
            l = mid
        } else {
            r = mid - 1
        }
    }

    return r
}
```


```go
// 浮点数二分模板
// check 检查x是否满足某种性质
func check(x float64) bool

func binarySearch(l, r float64) float64 {
    eps := 1e-6 // eps 表示精度，取决于题目对精度的要求, 如果是只取整数部分，eps = 1
    for (r - l) > eps {
        mid := (l + r) / 2
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }
    return l
}
```

相关LeetCode题:

> 整数二分
> 162.寻找峰值
> 
> 浮点数二分
> 69.求x的平方根


相关acwing题：
> 整数二分：
> 789. 数的范围 https://www.acwing.com/activity/content/code/content/41409

> 浮点数二分：
> 790. 数的三次方根 https://www.acwing.com/problem/content/792/

