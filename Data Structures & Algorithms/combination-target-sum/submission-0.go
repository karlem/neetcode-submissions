func combinationSum(nums []int, target int) [][]int {
    var current []int
	var out [][]int

	var backtrack func(int,int)
	backtrack = func(i int, currentSum int) {
		if currentSum > target {
			return
		}

		if currentSum == target {
			temp := make([]int, len(current));
			copy(temp, current);
			out = append(out, temp)
		}

		for j := i; j < len(nums); j++ {
			current = append(current, nums[j])
			backtrack(j, currentSum + nums[j])
			current = current[:len(current) - 1]
		}
	}

	backtrack(0, 0)
	return out
}
