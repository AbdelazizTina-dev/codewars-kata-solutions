package kata

func Partitions(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			dp[j] += dp[j-i]
		}
	}

	return dp[n]
}
