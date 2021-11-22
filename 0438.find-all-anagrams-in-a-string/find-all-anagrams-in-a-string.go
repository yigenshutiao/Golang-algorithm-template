package _438_find_all_anagrams_in_a_string

//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]

func findAnagrams(s string, p string) []int {
	pl := len(p)
	sl := len(s)
	if pl > sl {
		return nil
	}
	var result []int

	m := make(map[byte]int)
	for i := 0; i < pl; i++ {
		m[p[i]]++
	}

	for i1 := 0; i1 < pl; i1++ {
		m[s[i1]]--
	}

	for i := 0; i < sl-pl+1; i++ {

		flag := 0

		if i > 0 {
			m[s[i-1]]++
			m[s[i+pl-1]]--
		}

		for _, v := range m {
			if v != 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			result = append(result, i)
		}

	}
	return result
}
