func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i :=1 ; i < n ; i++ {
		n2, n3, n5 := dp[a] << 1, dp[b] * 3,  dp[c] * 5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			a ++
		}
		if dp[i] == n3 {
			b ++
		}
		if dp[i] == n5 {
			c ++
		}
	}

	return dp[n-1]
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}
	if y <= x && y <= z {
		return y
	}
	return z
}
