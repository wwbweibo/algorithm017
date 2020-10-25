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
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 824 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) (result int) {
	numRow := len(grid)
	if numRow <= 0 {
		return 0
	}

	numColumn := len(grid[0])

	for i :=0 ; i < numRow ;i ++ {
		for j := 0; j < numColumn; j ++ {
			if grid[i][j] == '1' {
				result += 1
				dfs(&grid, i, j)
			}
		}
	}

	return
}

func dfs(grid *[][]byte, row, column int ){
	(*grid)[row][column] = '0'
	numRow, numColumn := len(*grid), len((*grid)[0])
	for _, pos := range [][]int{{row - 1, column}, {row + 1, column}, {row, column - 1}, {row, column + 1}} {
		if 0 <= pos[0] && pos[0] < numRow && 0 <= pos[1] && pos[1] < numColumn && (*grid)[pos[0]][pos[1]] == '1' {
			dfs(grid, pos[0], pos[1])
		}
	}

}
//leetcode submit region end(Prohibit modification and deletion)
