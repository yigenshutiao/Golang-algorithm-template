package _049_group_anagrams

func groupAnagrams(strs []string) [][]string {
	var res [][]string

	kv := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, v := range s {
			cnt[v-'a'] += 1
		}

		kv[cnt] = append(kv[cnt], s)
	}

	for _, v := range kv {
		res = append(res, v)
	}

	return res
}
