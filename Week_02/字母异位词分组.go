func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	retmap := make(map[string][]string)
	for _, str := range strs {
		lst := make([]int, 26)
		ret := ""
		for _, char := range str {
			lst[char - 'a'] += 1
		}
		for _,count := range lst {
			ret += string(count) + ","
		}
		if arr,ok := retmap[ret]; ok {
			retmap[ret] = append(arr, str)
		} else {
			retmap[ret] = []string{str}
		}
	}
	for _,v := range retmap {
		result = append(result, v)
	}
	return result
}