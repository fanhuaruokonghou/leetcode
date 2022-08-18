package DynamicProgramming

func countBits(n int) []int {
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
