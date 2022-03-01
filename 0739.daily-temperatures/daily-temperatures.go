package _739_daily_temperatures

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))

	// 对于温度表里的每个温度
	for i := 0; i < len(temperatures); i++ {
		// 找到第一个大于t[i]的元素的索引值
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}
