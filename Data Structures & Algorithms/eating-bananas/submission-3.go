func getMax(nums []int) int {
	var maxN int
	for _, n := range nums {
		maxN = max(maxN, n)
	}
	return maxN
}

func totalEatingSpeed(piles []int, k int) int {
    var total int
    for _, pile := range piles {
        total += int(math.Ceil(float64(pile) / float64(k)))
    }
    return total
}

// 1, 2, 3, 4
func minEatingSpeed(piles []int, h int) int {
    left, right := 1, getMax(piles)

    for left < right {
        mid := int(left + (right-left) / 2)

        if totalEatingSpeed(piles, mid) <= h {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
