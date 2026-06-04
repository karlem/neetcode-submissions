type cords struct {
	i, j int
}

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	var bfs func(i, j int)
	bfs = func(i, j int) {
		q := []cords{{i:i, j:j}}
		var curr cords

		for len(q) > 0 {
			curr, q = q[0], q[1:]

			// look up
			if curr.i-1 >= 0 && grid[curr.i-1][curr.j] == '1' {
				q = append(q, cords{i: curr.i-1, j: curr.j})
			}

			// look down
			if curr.i+1 < rows && grid[curr.i+1][curr.j] == '1' {
				q = append(q, cords{i: curr.i+1, j: curr.j})
			}

			// look left
			if curr.j-1 >= 0 && grid[curr.i][curr.j-1] == '1' {
				q = append(q, cords{i: curr.i, j: curr.j-1})
			}

			// look right
			if curr.j+1 < cols && grid[curr.i][curr.j+1] == '1' {
				q = append(q, cords{i: curr.i, j: curr.j+1})
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