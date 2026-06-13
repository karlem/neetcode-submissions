func rob(nums []int) int {
	mem := map[int]int{}
    var rec func(i int) int
	rec = func(i int) int {
		if i >= len(nums) {
			return 0
		}

		if val, ok := mem[i]; ok {
			return val
		}

		mem[i] = max(nums[i]+rec(i+2), rec(i+1))
		return mem[i]
	}

	return max(rec(0), rec(1))
}
