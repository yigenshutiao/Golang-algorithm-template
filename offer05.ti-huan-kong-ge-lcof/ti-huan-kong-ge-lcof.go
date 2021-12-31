package offer05_ti_huan_kong_ge_lcof

func replaceSpace(s string) string {
	res := []byte(s)

	cnt := 0
	for i := 0; i < len(res); i++ {
		if res[i] == ' ' {
			cnt++
		}
	}

	tmp := make([]byte, cnt*2)
	res = append(res, tmp...)

	i := len(res) - 1
	j := len(s) - 1

	for j >= 0 {
		if res[j] != ' ' {
			res[i] = res[j]
			i--
			j--
		} else {
			res[i] = '0'
			res[i-1] = '2'
			res[i-2] = '%'
			j--
			i -= 3
		}
	}

	return string(res)
}
