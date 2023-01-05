package graph

// MaxAreaOfIsland
// @summary: LeetCode #695 (Medium)
// @author : binqibang
// @created: 2023/1/5 13:21
func MaxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j >= m || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 2
		return 1 + dfs(i, j-1) + dfs(i, j+1) + dfs(i-1, j) + dfs(i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}
	return maxArea
}
