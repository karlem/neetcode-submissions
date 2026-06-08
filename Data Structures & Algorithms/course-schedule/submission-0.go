func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, preq := range prerequisites {
		from, to := preq[0], preq[1]

		adj[from] = append(adj[from], to)
		inDegree[to]++
	}

	var q []int

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var takenCourses int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		takenCourses++

		for _, neigbor := range adj[node] {
			inDegree[neigbor]--
			if inDegree[neigbor] == 0 {
				q = append(q, neigbor)
			}
		}
	}

	return takenCourses == numCourses
}
