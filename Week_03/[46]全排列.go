//给定一个 没有重复 数字的序列，返回其所有可能的全排列。 
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法 
// 👍 956 👎 0

package main


//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(int)
	backtrack = func(curr int) {
		if curr == len(nums) {
			ans := make([]int, len(nums))
			copy(ans, nums)
			result = append(result,ans)
		}

		for i := curr ;i < len(nums) ; i ++ {
			nums[curr], nums[i] = nums[i], nums[curr]
			backtrack(curr + 1)
			nums[curr], nums[i] = nums[i], nums[curr]
		}

	}

	backtrack(0)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
