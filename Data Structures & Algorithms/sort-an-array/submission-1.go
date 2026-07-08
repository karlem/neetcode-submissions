func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	half := len(nums)/2
	left, right := sortArray(nums[:half]), sortArray(nums[half:])
	i, j := 0,0
	
	res := []int{}
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
