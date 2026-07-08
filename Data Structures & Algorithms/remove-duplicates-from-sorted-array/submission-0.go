// [1,2,3,4,1,1,,,]

func removeDuplicates(nums []int) int {
	k := 0
	prev := 101

	for i := range nums {
		if nums[i] != prev {
			prev = nums[i]
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
