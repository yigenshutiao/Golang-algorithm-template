package offer21_diao_zheng_shu_zu_shun_xu_shi_qi_shu_wei_yu_ou_shu_qian_mian_lcof

func exchange(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j]%2 == 0 && nums[j+1]%2 != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
