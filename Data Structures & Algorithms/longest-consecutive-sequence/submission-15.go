func longestConsecutive(nums []int) int {
	sequence := map[int]struct{}{}

	for _, num := range nums {
		sequence[num] = struct{}{}
	}

	best := 0
	for _, num := range nums {
		if _, ok := sequence[num-1]; ok {
			continue
		}

		l := 1
		for {
			if _, ok := sequence[num+l]; !ok {
				break
			}

			l++
		}

		best = max(best, l)
	}

	return best
}
