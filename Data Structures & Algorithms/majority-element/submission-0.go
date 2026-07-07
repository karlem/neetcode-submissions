// [5,5,1,1,1,5,5]

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}

		if count == 0 {
			candidate = nums[i]
			count++
		}
	}

	return candidate
}
