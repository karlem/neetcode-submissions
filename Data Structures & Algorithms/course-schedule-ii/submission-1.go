func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := map[int][]int{}
	deps := make([]int, numCourses)

	for _, preq := range prerequisites {
		from, to := preq[0], preq[1]
		deps[from]++
		adj[to] = append(adj[to], from)
	}	

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if deps[i] == 0 {
			q = append(q, i)
		}
	}

	var out []int
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]

		out = append(out, pop)

		for _, neighbor := range adj[pop] {
			deps[neighbor]--
			if deps[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}
	

	if len(out) != numCourses {
		return []int{}
	}

	return out	
}
