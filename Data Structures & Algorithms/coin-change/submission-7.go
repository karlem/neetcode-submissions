func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if a-coin >= 0 {
				dp[a] = min(dp[a], dp[a-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
        return -1
    }

	return dp[amount]
}