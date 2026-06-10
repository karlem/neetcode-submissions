func combinationSum(nums []int, target int) [][]int {
    var res [][]int
	var curr []int
	sort.Ints(nums)

	var dfs func(int, int)
	dfs = func(i, currSum int) {
		if currSum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for j := i; j < len(nums); j++ {
			if currSum + nums[j] > target {
				return
			}

			curr = append(curr, nums[j])
			dfs(j, currSum + nums[j])
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, 0)
	return res
}
