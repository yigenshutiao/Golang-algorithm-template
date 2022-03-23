package _99_evaluate_division

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}
	for i := 0; i < len(equations); i++ { //建立图
		a, b := equations[i][0], equations[i][1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1 / values[i]
		graph[a][a] = 1.0
		graph[b][b] = 1.0
	}
	var ans []float64
	for _, query := range queries {
		a, b := query[0], query[1]
		ans = append(ans, calc(graph, a, b, map[string]bool{}))
	}
	return ans
}

func calc(graph map[string]map[string]float64, a, b string, visit map[string]bool) float64 {
	if _, ok := graph[a]; !ok { //没有a字符
		return -1.0
	}
	if _, ok := graph[b]; !ok { //没有b字符
		return -1.0
	}
	if val, ok := graph[a][b]; ok { //有a / b的值
		return val
	} else { // 没有a / b 的值，则遍历a的所有临接点，并记录已经访问过a字符，防止重复遍历
		visit[a] = true
		for key, val := range graph[a] {
			if visit[key] {
				continue
			}
			v := calc(graph, key, b, visit)
			if v != -1.0 {
				return val * v
			}
		}
	}
	return -1.0
}
