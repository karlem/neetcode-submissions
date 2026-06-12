func canJump(nums []int) bool {
    var rec func(int) bool
	rec = func(i int) bool {
		if i == len(nums)-1 {
			return true
		}

		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums) {
				break
			}

			if rec(i+j) {
				return true
			}
		}

		return false
	}

	return rec(0)
}
