package DynamicProgramming

func GetRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	dp[0] = make([]int, 1)
	dp[0][0] = 1

	for i := 1; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i>>1+1; j++ {
			if i == 1 {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			dp[i][i-j] = dp[i][j]
		}
	}
	return dp[rowIndex]
}
