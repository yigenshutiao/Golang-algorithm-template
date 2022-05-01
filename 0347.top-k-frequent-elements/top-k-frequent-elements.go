package _347_top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	kv := make(map[int]int)

	for _, v := range nums {
		kv[v]++
	}

	var a []int
	for k := range kv {
		a = append(a, k)
	}

	sort.Slice(a, func(i, j int) bool {
		// 按照kv对中k出现的次数降序排序
		return kv[a[i]] > kv[a[j]]
	})

	return a[:k]
}
