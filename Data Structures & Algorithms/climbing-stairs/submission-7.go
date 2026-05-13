// top down
func climbStairs(n int) int {
	mem := map[int]int{}

	var dfs func(i int) int 
    dfs = func(i int) int {
		if i <= 2 {
			return i
		}

		if v, ok := mem[i]; ok {
			return v
		}

		mem[i] = dfs(i-1) + dfs(i-2)
		return mem[i]
	}

	return dfs(n)
}
