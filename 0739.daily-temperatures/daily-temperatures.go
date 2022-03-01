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

func dailyTemperature(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int

	for i := 0; i < len(temperatures); i++ {
		// 当前元素
		cur := temperatures[i]
		// 当前元素大于栈顶元素
		for len(stack) > 0 && cur > temperatures[stack[len(stack)-1]] {
			// 获取栈顶元素的下标
			idx := stack[len(stack)-1]
			// 将栈顶元素出栈
			stack = stack[:len(stack)-1]
			// 更新结果表
			res[idx] = i - idx
		}
		// 将新元素入栈
		stack = append(stack, i)
	}
	return res
}
