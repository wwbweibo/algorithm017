//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
// 
// Related Topics 字符串 回溯算法 
// 👍 1415 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backTrack("", 0, 0, n, &ans)
	return ans
}

func backTrack(s string, left, right, n int, ans *[]string)  {
	if len(s) == 2 * n {
		*ans = append(*ans, s)
		return
	}
	if left < n {
		s = s + "("
		backTrack(s, left + 1, right, n, ans)
		s = s[0 : len(s) - 1]
	}
	if right < left {
		s = s + ")"
		backTrack(s, left, right + 1, n , ans)
		s = s[0 : len(s) - 1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
