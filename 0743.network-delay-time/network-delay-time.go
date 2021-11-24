package _743_network_delay_time

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	// distance: 源点k到各点的最短距离
	distance := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distance[i] = math.MaxInt32
	}
	// 点到点之间的距离, ---- ！注意这里需要初始化一下，不能为默认0，因为输入可能是0
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			graph[i][j] = -1
		}
	}
	for i := 0; i < len(times); i++ {
		s, e, v := times[i][0], times[i][1], times[i][2]
		graph[s][e] = v
		if s == k {
			distance[e] = v
		}
	}
	//fmt.Println(graph)

	selected := make(map[int]int, n+1)
	selected[k] = 1

	for j := 1; j <= n; j++ {
		// 从U中（unvisited）选择距离源点最近的点  ---- 可以优化(小顶堆，优先级队列)
		min, minDot := math.MaxInt32, 0
		for i := 1; i <= n; i++ {
			if selected[i] == 0 && distance[i] < min {
				min = distance[i]
				minDot = i
			}
		}
		selected[minDot] = 1
		if minDot <= 0 {
			break
		}

		// 对于minDot的所有邻接点v，对比distance[v] 和 distance[minDot] + grahp[minDot][v] 刷新距离
		for i := 1; i <= n; i++ {
			if graph[minDot][i] >= 0 {
				if distance[minDot]+graph[minDot][i] < distance[i] {
					distance[i] = distance[minDot] + graph[minDot][i]
				}
			}
		}
		//fmt.Println("chose ", minDot,  distance)
	}

	max := 0
	for i, v := range distance {
		if v > max && i != k {
			max = v
		}
	}

	if max == math.MaxInt32 {
		return -1
	}
	return max
}
