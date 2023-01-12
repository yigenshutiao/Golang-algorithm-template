package offer04_er_wei_shu_zu_zhong_de_cha_zhao_lcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}
	for m, n := 0, len(matrix[0])-1; m < len(matrix) && n >= 0; {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			n--
		} else if matrix[m][n] < target {
			m++
		}
	}
	return false
}
