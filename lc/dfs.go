package lc

//主要实现dfs深度优先遍历 等相关案例

// lc-200 岛屿数量
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	result := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '2'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				result++
				dfs(row, col)
			}
		}
	}
	return result
}
