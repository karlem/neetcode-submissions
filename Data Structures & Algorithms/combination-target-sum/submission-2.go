func combinationSum(nums []int, target int) [][]int {
    var current []int
	var out [][]int
	sort.Ints(nums)

	var backtrack func(int,int)
	backtrack = func(i int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(current));
			copy(temp, current);
			out = append(out, temp)
			return
		}

		for j := i; j < len(nums); j++ {
			if currentSum + nums[j] > target {
                return
            }

			current = append(current, nums[j])
			backtrack(j, currentSum + nums[j])
			current = current[:len(current) - 1]
		}
	}

	backtrack(0, 0)
	return out
}
