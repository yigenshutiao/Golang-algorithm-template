package _034_find_first_and_last_position_of_element_in_sorted_array

// 二分模板 1：
// 当我们将区间[l, r]划分成[l, mid]和[mid + 1, r]时，其更新操作是r = mid或者l = mid + 1,计算mid时不需要加1，即mid = (l + r)/2。

// int bsearch_1(int l, int r)
//{
//    while (l < r)
//    {
//        int mid = (l + r)/2;
//        if (check(mid)) r = mid;
//        else l = mid + 1;
//    }
//    return l;
//}

// 二分模板 2：
//  当我们将区间[l, r]划分成[l, mid - 1]和[mid, r]时，其更新操作是r = mid - 11或者l = mid，此时为了防止死循环，计算mid时需要加1，即mid = ( l + r + 1 ) /2。

//int bsearch_2(int l, int r)
//{
//    while (l < r)
//    {
//        int mid = ( l + r + 1 ) /2;
//        if (check(mid)) l = mid;
//        else r = mid - 1;
//    }
//    return l;
//}

// 区间范围：
// 左起点：第一个 >= target 的值
// 右终点：最后一个 <= target 的值
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return []int{-1, -1}
	}

	start := left

	left, right = 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	end := right

	return []int{start, end}
}
