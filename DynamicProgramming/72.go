package DynamicProgramming

func climbStairs72(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		if i == 1 {
			dp[i] = 1
			continue
		}
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs72Method2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i < n+1; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
