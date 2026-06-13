func minCostClimbingStairs(cost []int) int {
	mem := map[int]int{}

    var rec func(i int) int
	rec = func(i int) int {
		if i >= len(cost) {
			return 0
		}

		if val, ok := mem[i]; ok {
			return val
		}

		mem[i] = cost[i] + min(rec(i+1), rec(i+2))
		return mem[i]
	}

	return min(rec(0), rec(1))
}
