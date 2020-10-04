func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make( map[rune]int)
	for _,char := range s {
			set[char] = set[char] + 1
	}

	for _, char := range t {
		if _, ok := set[char]; ok {
			set[char] = set[char] - 1
			if set[char] < 0 {
				return false
			}
		}else {
			return false
		}
	}

	return true
}