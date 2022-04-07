package _049_group_anagrams

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	// 这个特殊的数据结构需要记忆一下...
	wordCnt := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, v := range s {
			cnt[v-'a']++
		}
		wordCnt[cnt] = append(wordCnt[cnt], s)
	}

	for _, words := range wordCnt {
		res = append(res, words)
	}

	return res
}
