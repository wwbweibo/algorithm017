//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性： 
//
// 
// 每行中的整数从左到右按升序排列。 
// 每行的第一个整数大于前一行的最后一个整数。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 3
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 13
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：matrix = [], target = 0
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 0 <= m, n <= 100 
// -104 <= matrix[i][j], target <= 104 
// 
// Related Topics 数组 二分查找 
// 👍 264 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}
	for _, row := range matrix {
		if len(row) <= 0 {
			return false
		}
		if row[0] <= target && target <= row[len(row) - 1] {
			return binarySearch(row, target) != -1
		}
	}
	return false
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid]	 == target {
			return target
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
