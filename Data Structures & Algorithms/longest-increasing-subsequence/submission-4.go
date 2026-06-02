func lengthOfLIS(nums []int) int {
    // for each i in nums I need to try the path forward and count how many of the elements j
	// are bigger then the previous one.
	// i = 0
	// j = 1 for each j = i+1...len(nums); check if nums[i] < nums[j] and if yes
	// then check from there by recursively calling in with rec(j).
	// so: if nums[i] < nums[j] then currMax = max(currMax, rec(j) + 1)
	// the base case if the end of the array nums in that case we say it is 1 because it is 1
	// number in subsequence.
	mem := map[int]int{}

	var rec func(int) int
	rec = func(i int) int {
		if i == len(nums)-1 {
			return 1
		}

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
