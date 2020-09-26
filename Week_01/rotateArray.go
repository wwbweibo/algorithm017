func rotate(nums []int, k int)  {
	length := len(nums)
	k = k % length
	count := 0
	for  start:= 0; count < length; start ++ {
		current := start
		prev := nums[start]
		flag := true

		for ; flag; {
			next := (current + k) % length
			temp := nums[next]
			nums[next] = prev
			current = next
			prev = temp
			count ++
			flag = current != start
		}
	}
}