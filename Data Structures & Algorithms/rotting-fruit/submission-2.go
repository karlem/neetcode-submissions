func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])

	var fresh int
	q := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var minutes int
	for fresh > 0 && len(q) > 0 {
		size := len(q)
		
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]

			for _, dir := range directions {
				r, c := pop[0]+dir[0], pop[1]+dir[1]

				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					q = append(q, [2]int{r, c})
					fresh--
					grid[r][c] = 2
				}
			}
		}

		minutes++
	}

	if fresh != 0 {
		return -1
	}

	return minutes
}
