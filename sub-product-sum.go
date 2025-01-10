package kata

const MOD = 1000000007

func ProductSum(a []int, m int) int {
	n := len(a)

	// DP table to store the sum of products of subsequences of length k
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base case: there's one subsequence of length 0 (the empty subsequence)
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}

	// Fill the DP table
	for j := 1; j <= n; j++ {
		for k := 1; k <= m; k++ {
			dp[k][j] = dp[k][j-1]                             // Case when we don't include a[j-1]
			dp[k][j] = (dp[k][j] + dp[k-1][j-1]*a[j-1]) % MOD // Case when we include a[j-1]
		}
	}

	// The answer is the sum of all subsequences of length m
	return dp[m][n]
}
