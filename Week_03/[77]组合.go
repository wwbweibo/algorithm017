//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法 
// 👍 416 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans :=[][]int{}
	temp := []int{}

	var generator func(int)
	generator = func(current int) {
		if len(temp) + (n - current + 1) < k {
			return
		}
		if len(temp) == k {
			a := make([]int, k)
			copy(a, temp)
			ans = append(ans, a)
			return
		}

		temp = append(temp, current)
		generator(current + 1)
		temp = temp[0 : len(temp) -1]

		generator(current + 1)
	}

	generator(1)

	return ans
	
}
//leetcode submit region end(Prohibit modification and deletion)
