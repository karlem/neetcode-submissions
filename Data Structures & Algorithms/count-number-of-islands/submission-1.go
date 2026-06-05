type cords struct {
	i, j int
}

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	directions := [][2]int{
		{-1, 0},
		{1,0},
		{0, -1},
		{0, 1},
	}

	var bfs func(i, j int)
	bfs = func(i, j int) {
		q := []cords{{i:i, j:j}}
		var curr cords

		
		for len(q) > 0 {
			curr, q = q[0], q[1:]

			for _, d := range directions {
				r, c := d[0]+curr.i, d[1]+curr.j

				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == '1' {
					q = append(q, cords{i: r, j: c})
				}
			}

			grid[curr.i][curr.j] = '0'
		}
	}
	
	var islands int
	for i := range rows {
		for j := range cols {
			if grid[i][j] == '1' {
				bfs(i, j)
				islands++
			}			
		}
	}

	return islands
}