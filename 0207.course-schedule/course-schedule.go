package _207_course_schedule

func canFinish(n int, edges [][]int) bool {
	g := make([][]int, n)
	d := make([]int, n)

	for _, e := range edges {
		b, a := e[0], e[1]
		g[a] = append(g[a], b)
		d[b]++
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if d[i] == 0 {
			q = append(q, i)
		}
	}

	cnt := 0
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		cnt++

		for _, i := range g[t] {
			d[i]--
			if d[i] == 0 {
				q = append(q, i)
			}
		}
	}

	return cnt == n
}
