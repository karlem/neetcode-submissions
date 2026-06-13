func permute(nums []int) [][]int {
	var curr []int
	var result [][]int
	used := map[int]bool{}

	var dfs func()
	dfs = func() {
		if len(curr) == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}

		for _, num := range nums {
			if used[num] {
				continue
			}

			curr = append(curr, num)
			used[num] = true
			dfs()
			used[num] = false
			curr = curr[:len(curr)-1]
		}
	}

	dfs()
	return result
}
