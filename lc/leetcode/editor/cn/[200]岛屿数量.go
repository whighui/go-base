//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] 的值为 '0' 或 '1' 
// 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2270 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)

// lc-200 岛屿数量 就是岛屿数量呗
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	result := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		//这里边递归退出条件就是一定要记住是 grid[i][j]!='1' 因为下方存在设置为2了呢
		if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != '1' {
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

//leetcode submit region end(Prohibit modification and deletion)
