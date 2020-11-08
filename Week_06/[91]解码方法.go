//一条包含字母 A-Z 的消息通过以下方式进行了编码： 
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
// 
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。 
//
// 题目数据保证答案肯定是一个 32 位的整数。 
//
// 
//
// 示例 1： 
//
// 输入："12"
//输出：2
//解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
// 
//
// 示例 2： 
//
// 输入："226"
//输出：3
//解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
// 
//
// 示例 3： 
//
// 输入：s = "0"
//输出：0
// 
//
// 示例 4： 
//
// 输入：s = "1"
//输出：1
// 
//
// 示例 5： 
//
// 输入：s = "2"
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 100 
// s 只包含数字，并且可以包含前导零。 
// 
// Related Topics 字符串 动态规划 
// 👍 552 👎 0

package main

import "strconv"
//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[0] = 1
	dp[1] = 1
	if len(s) == 1 {
		return dp[1]
	}

	for i:= 2; i <= len(s) ; i ++ {
		num, _ := strconv.Atoi(string(s[i-1]))
		num2, _ := strconv.Atoi(string(s[i-2]))
		if num  + num2 == 0 || (num ==0 && num2 > 2){
			return 0
		} else if num == 0{
			dp[i] = dp[i - 2]
		} else if num2 == 0 {
			dp[i] = dp[i - 1]
 		} else if  num2 * 10  + num > 26 {
 			dp[i] = dp[i - 1]
		} else {
			dp[i] = dp[i - 2] + dp[i-1]
		}
	}
	return dp[len(s)]
}
//leetcode submit region end(Prohibit modification and deletion)