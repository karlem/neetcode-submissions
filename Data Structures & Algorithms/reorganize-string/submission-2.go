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
	// max heap
	// with queue the queue get moved back to heap once the heap is empty
	// this could be slice of len 26
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for value, count := range freq {
		heap.Push(h, Item{
			value,
			count,
		})
	}

	res := []rune{}
	var prev *Item

	for h.Len() > 0 || prev != nil {
		 if prev != nil && h.Len() == 0 {
            return ""
        }
		
		item := heap.Pop(h).(Item)
		res = append(res, item.value)
		item.count--;
		
		if prev != nil {
			heap.Push(h, *prev)
			prev = nil
		}

		if item.count > 0 {
			prev = &item
		}
	}

	return string(res)
}