func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := map[int][]int{}
	inDegree := make([]int, numCourses)

	for _, preq := range prerequisites {
		from, to := preq[0], preq[1]
		adj[from] = append(adj[from], to)
		inDegree[to]++
	}
	
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var seen int
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		seen++

		for _, neighbor := range adj[pop] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
		
	}

	return seen == numCourses
}
