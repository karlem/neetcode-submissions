func maxSubArray(nums []int) int {
	// curr = 0
	// best = 0
    // curr = max(curr + nums[i], nums[i])
	// best = max(best, curr)

	curr := 0
	best := -1
	for _, num := range nums {
		curr = max(curr+num, num)
		best = max(best, curr)
	}

	return best
}
