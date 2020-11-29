//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。 
//
// 
// 如果剩余字符少于 k 个，则将剩余字符全部反转。 
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。 
// 
//
// 
//
// 示例: 
//
// 输入: s = "abcdefg", k = 2
//输出: "bacdfeg"
// 
//
// 
//
// 提示： 
//
// 
// 该字符串只包含小写英文字母。 
// 给定字符串的长度和 k 在 [1, 10000] 范围内。 
// 
// Related Topics 字符串 
// 👍 105 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	str := []byte(s)
	for start := 0; start < len(s); start = start + (2 * k) {
		i, j := start, start+k-1
		if j >= len(s) {
			j = len(s) - 1
		}
		for i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}

	return string(str)
}

//leetcode submit region end(Prohibit modification and deletion)
