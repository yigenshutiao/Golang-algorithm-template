package _088_merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j++
	}

	quickSort(nums1)
}

func quickSort(a []int) {
	if len(a) < 1 {
		return
	}

	idx, pos := 0, len(a)-1
	for i := len(a) - 1; i >= 1; i-- {
		if a[i] > a[idx] {
			a[i], a[pos] = a[pos], a[i]
			pos--
		}
	}

	a[idx], a[pos] = a[pos], a[idx]
	quickSort(a[:pos])
	quickSort(a[pos+1:])
}
