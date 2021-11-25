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

func networkDelayTimes(times [][]int, n, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x][y] = t[2]
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range g[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}

	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
