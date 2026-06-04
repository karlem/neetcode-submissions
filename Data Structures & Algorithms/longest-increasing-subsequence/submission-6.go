func lengthOfLIS(nums []int) int {
	mem := map[int]int{}
	
    var rec func(int) int
	rec = func(i int) int {
		if val, ok := mem[i]; ok {
			return val
		}

		best := 1
		for j := i+1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				best = max(best, rec(j) + 1)
			}
		}

		mem[i] = best
		return best
	}

	best := 1
	for i := range nums {
		best = max(best, rec(i))
	}

	return best
}
