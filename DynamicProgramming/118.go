package DynamicProgramming

func generate118(numRows int) [][]int {
	dp := make([][]int, numRows)
	dp[0] = make([]int, 1)
	dp[0][0] = 1
	for i := 1; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j < i>>1+1; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
			dp[i][i-j] = dp[i][j]
		}
	}
	return dp
}
