package DynamicProgramming

func countBits338(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		j := i
		for j > 0 {
			if j&1 > 0 {
				dp[i]++
			}
			j = j >> 1
		}
	}
	return dp
}

func countBits338Method2(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}
