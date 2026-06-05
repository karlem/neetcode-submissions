func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	fresh := 0
	q := [][2]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	minutes := 0

	for fresh > 0 && len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			i, j := q[0][0], q[0][1]
			q = q[1:]

			for _, d := range directions {
				r, c := d[0]+i, d[1]+j
				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					fresh--
					q = append(q, [2]int{r, c})
				}
			}
		}

		minutes++
	}

	if fresh == 0 {
		return minutes
	}

	return -1
}
