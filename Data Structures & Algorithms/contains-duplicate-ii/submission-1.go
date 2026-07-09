func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]struct{}{}

	l := 0
	for	r := 0; r < len(nums); r++ {
		for r-l > k {
			delete(m, nums[l])
			l++
		}

		if _, ok := m[nums[r]]; ok {
			return true
		}
		
		m[nums[r]] = struct{}{}
	}

	return false
}
