//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。 
//
// 说明：每次只能向下或者向右移动一步。 
//
// 示例: 
//
// 输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
// 
// Related Topics 数组 动态规划 
// 👍 712 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	for i := row - 2; i >= 0 ; i -- {
		grid[i][col - 1] += grid[i  + 1][col-1]
	}

	for j := col - 2;  j >= 0; j -- {
		grid[row - 1][j] += grid[row-1][j + 1]
	}

	for i := row - 2; i >= 0; i -- {
		for j := col - 2; j >= 0; j -- {
			grid[i][j] += min(grid[i + 1][j] , grid[i][j + 1])
		}
	}

	return grid[0][0]
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
