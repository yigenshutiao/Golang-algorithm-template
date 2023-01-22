package offer15_er_jin_zhi_zhong_1de_ge_shu_lcof

func hammingWeight(num uint32) int {
	res := 0

	for num != 0 {
		tmp := num & 1
		res += int(tmp)
		num = num >> 1
	}

	return res
}
