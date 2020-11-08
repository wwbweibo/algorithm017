//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。 
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2： 
//
// 
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
//
// 
//
// 提示： 
//
// 
// 0 <= word1.length, word2.length <= 500 
// word1 和 word2 由小写英文字母组成 
// 
// Related Topics 字符串 动态规划 
// 👍 1239 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	if n * m == 0 {
		return n + m
	}

	dp := make([][]int, n + 1)
	for i := 0 ; i <= n; i ++ {
		dp[i] = make([]int, m + 1)
	}

	for i := 0; i <= n; i ++ {
		dp[i][0] = i
	}

	for j := 0; j <= m; j ++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i ++ {
		for j := 1; j <= m; j ++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1]{
				left_down += 1
			}
			dp[i][j] = min(left, down, left_down)
		}
	}
	return dp[n][m]
}

func min(x, y , z int) int {
	if x < y && x < z {
		return x
	}
	if y < x && y < z {
		return y
	}
	return z
}
//leetcode submit region end(Prohibit modification and deletion)
