import (
	"slices"
	"cmp"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return cmp.Compare(a[0], b[0]) })

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		prevIndex := len(merged)-1
		prev := merged[prevIndex]

		if intervals[i][0] <= prev[1] {
			merged[prevIndex][1] = max(prev[1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
