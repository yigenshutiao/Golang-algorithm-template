package _017_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) < 1 {
		return res
	}
	kv := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var dfs func(res *[]string, i int, cur string)
	dfs = func(res *[]string, i int, cur string) {
		if len(cur) == len(digits) {
			*res = append(*res, cur)
			return
		}

		// 单层逻辑
		k := string(digits[i])
		for _, v := range kv[k] {
			cur = cur + string(v)
			dfs(res, i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(&res, 0, "")
	return res
}
