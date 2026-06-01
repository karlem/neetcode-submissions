func lengthOfLIS(nums []int) int {
	mem := map[int]int{}

	var rec func(int) int
	rec = func(i int) int {
		if i == len(nums) - 1 {
			return 1
		}

		if val, ok := mem[i]; ok {
			return val
		}

		best := 1

		for j := i+1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				best = max(best, rec(j) + 1)
			}
		}

		mem[i] = best
		return mem[i]
	}

	overallBest := 0
	for i := 0; i < len(nums); i++ {
		overallBest = max(overallBest, rec(i))
	}
	
	return overallBest
}
