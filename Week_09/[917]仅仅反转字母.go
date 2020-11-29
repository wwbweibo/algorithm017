//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入："ab-cd"
//输出："dc-ba"
// 
//
// 示例 2： 
//
// 输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
// 
//
// 示例 3： 
//
// 输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
// 
//
// 
//
// 提示： 
//
// 
// S.length <= 100 
// 33 <= S[i].ASCIIcode <= 122 
// S 中不包含 \ or " 
// 
// Related Topics 字符串 
// 👍 67 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(S string) string {
	result := make([]byte,0 )
	j := len(S) - 1
	for i:=0; i < len(S); i ++ {
		if isChar(S[i]) {
			for !isChar(S[j]) {
				j --
			}
			result = append(result, S[j])
			j --
		}else {
			result = append(result, S[i])
		}
	}

	return string(result)
}
func isChar(char uint8) bool {
	if 'a' <= char && char <= 'z' {
		return true
	} else if 'A' <= char && char <= 'Z' {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
