func maxArea(heights []int) int {
	best := 0
	left, right := 0, len(heights) - 1

	for left < right {
		best = max(best, (right - left) * min(heights[left], heights[right]))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	
	return best
}
