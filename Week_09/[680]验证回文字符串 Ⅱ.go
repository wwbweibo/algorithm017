//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。 
//
// 示例 1: 
//
// 
//输入: "aba"
//输出: True
// 
//
// 示例 2: 
//
// 
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
// 
//
// 注意: 
//
// 
// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。 
// 
// Related Topics 字符串 
// 👍 288 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	var validator func(low, high int, flag bool) bool

	validator = func(low, high int, flag bool) bool {
		for low < high {
			if s[low] == s[high] {
				low++
				high--
			} else if flag {
				return validator(low+1, high, false) || validator(low, high-1, false)
			} else {
				return false
			}
		}
		return true
	}

	return validator(0, len(s)-1, true)
}

//leetcode submit region end(Prohibit modification and deletion)
