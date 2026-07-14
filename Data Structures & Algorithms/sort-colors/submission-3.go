func sortColors(nums []int) {
	r, b := 0, len(nums) - 1
	w := 0
	for w <= b {
		if nums[w] == 0 {
			nums[r], nums[w] = nums[w], nums[r]
			r++
			w++
		} else if nums[w] == 2 {
			nums[b], nums[w] = nums[w], nums[b]
			b--
		} else {
			w++
		}
	}
}
