func removeElement(nums []int, val int) int {
    w := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[w] = nums[i]
			w++
		}
	}
	return w
}
