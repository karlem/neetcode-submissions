// [1,2,3,4,5,6,7,8]
// 8,7,6,5,4,3,2,1
// 5,6,7,8,4,3,2,1
// 5,6,7,8,1,2,3,4

func rotate(nums []int, k int) {
	l := len(nums)
	r := k%l

	reverse(nums, 0, l-1)
	reverse(nums, 0, r-1)
	reverse(nums, r, l-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}