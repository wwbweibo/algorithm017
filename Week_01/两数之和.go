func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i := 0; i < len(nums); i ++ {
		t := target - nums[i]
		idx, ok := set[t]
		if ok {
			return []int{idx, i}
		} else {
			set[nums[i]] = i
		}
	}
	return []int{}
}