func plusOne(digits []int) []int {
	flag := true
	for i := len(digits) - 1; i >= 0; i -- {
		// 需要加一，并且加一之后小于10，只需要将该位直接加一
		if flag && digits[i] + 1 < 10 {
			digits[i] += 1
			flag = false
		} else if flag && digits[i] + 1 >= 10 {
			digits[i] = 0
			flag = true
		}
	}

	if flag {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}