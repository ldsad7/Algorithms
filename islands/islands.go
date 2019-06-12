package islands

func Color(grid [][]uint8, i, j int) {
	grid[i][j] = 0
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		Color(grid, i+1, j)
	}
	if i-1 > -1 && grid[i-1][j] == 1 {
		Color(grid, i-1, j)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
		Color(grid, i, j+1)
	}
	if j-1 > -1 && grid[i][j-1] == 1 {
		Color(grid, i, j-1)
	}
}

// CountIslands returns the number of islands on a 2D grid map
func CountIslands(grid [][]uint8) int {
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				Color(grid, i, j)
				islands++
			}
		}
	}
	return islands
}
