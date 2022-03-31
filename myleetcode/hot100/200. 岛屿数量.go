package main

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}
