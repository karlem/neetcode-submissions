/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import (
	"slices"
	"cmp"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h MinHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *MinHeap) Peak() any {
	s := *h
	l := len(s)-1

	return s[l]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	l := len(old)-1
	pop := old[l]
	old = old[:l]
	*h = old
	return pop
}

func minMeetingRooms(intervals []Interval) int {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	h := &MinHeap{}
	heap.Init(h)

	for _, interval := range intervals {
		hp := h.Len()
		if hp > 0 && interval.start >= (*h)[0] {
			heap.Pop(h)
		}

		heap.Push(h, interval.end)
	}
	return h.Len()
}
