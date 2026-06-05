type Item struct {
	value rune
	count int
} 

type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reorganizeString(s string) string {
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for v, c := range freq {
		heap.Push(h, Item{value: v, count: c})
	}

	out := []rune{}
	var prev *Item

	for h.Len() > 0 {
		var current *Item
		if h.Len() > 0 {
			curr := heap.Pop(h).(Item)
			out = append(out, curr.value)
			curr.count--
			current = &curr
		}

		if prev != nil {
			heap.Push(h, *prev)
			prev = nil
		}

		if current != nil && current.count > 0 {
			prev = current
		}
	}

	if prev != nil {
		return ""
	}

	return string(out)
}
