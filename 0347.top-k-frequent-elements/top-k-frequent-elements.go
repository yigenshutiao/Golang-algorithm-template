package _347_top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	kv := map[int]int{}
	for _, num := range nums {
		kv[num] += 1
	}

	var a []int
	for k, _ := range kv {
		a = append(a, k)
	}

	// 按照出现的频率降序排序
	sort.Slice(a, func(i, j int) bool {
		return kv[a[i]] > kv[a[j]]
	})
	return a[:k]
}
