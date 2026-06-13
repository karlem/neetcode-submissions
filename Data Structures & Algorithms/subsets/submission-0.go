func subsets(nums []int) [][]int {
	var res [][]int
	var curr []int

	var dfs func(i int)
	dfs = func(i int) {
		cp := make([]int, len(curr))
		copy(cp, curr)
		res = append(res, cp)

		for j := i; j < len(nums); j++ {
			curr = append(curr, nums[j])
			dfs(j+1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)
	return res
}
