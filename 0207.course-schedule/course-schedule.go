package _207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录每个元素的出度
	graph := make([][]int, numCourses)
	// 记录从 0 到 n-1，所有元素的入度
	degreeIn := make([]int, numCourses)

	for _, pair := range prerequisites {
		cur, rely := pair[0], pair[1]
		degreeIn[cur]++
		graph[rely] = append(graph[rely], cur)
	}

	queue := make([]int, 0)
	// 把所有入度为0的元素放在queue里面
	for node, val := range degreeIn {
		if val == 0 {
			queue = append(queue, node)
		}
	}

	var cnt int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		cnt++

		// 遍历元素所有的出度，把出度元素的入度--
		for _, val := range graph[node] {
			degreeIn[val]--
			if degreeIn[val] == 0 {
				queue = append(queue, val)
			}
		}
	}

	return cnt == numCourses
}
