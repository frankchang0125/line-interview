package board_game

// For dice with 6 faces:
// 	dp[i] = dp[i-1] + dp[i-2] + dp[i-3] + dp[i-4] + dp[i-5] + dp[i-6]
func boardGame(target int) uint64 {
	dp := make([]uint64, target+1)

	// Only one way to reach '0' point.
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := 1; j <= 6; j++ {
			// Watch out for: dp[1] ~ dp[5]
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[target]
}
