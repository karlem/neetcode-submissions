// bottom up
func climbStairs(n int) int {
    if n <= 2 {
		return n
	}
    steps := make([]int, n)
	steps[0] = 1
	steps[1] = 2

	for i := 2; i < n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}

	return steps[n-1]
}
