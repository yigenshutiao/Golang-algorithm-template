package _017_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	var res []string
	if digits == "" {
		return res
	}

	digMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(curIdx int, cur string, res *[]string)
	dfs = func(curIdx int, cur string, res *[]string) {
		if curIdx == len(digits) {
			*res = append(*res, cur)
			return
		}

		for _, v := range digMap[digits[curIdx]] {
			cur = cur + string(v)
			dfs(curIdx+1, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, "", &res)

	return res
}
