type Item struct {
	Value byte
	Priority int
}

type PriorityHeap []Item

func (h PriorityHeap) Len() int           { return len(h) }
func (h PriorityHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority }
func (h PriorityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *PriorityHeap) Pop() any {
	old := *h
	l := len(old)
	item := old[l-1]
	*h = old[0 : l-1]
	return item
}

type QueueItem struct {
	HeapItem Item
	CooldDownEnd int
}

func leastInterval(tasks []byte, n int) int {
	freq := map[byte]int{}
	h := &PriorityHeap{}
	q := []QueueItem{}
	out := []byte{}

	for _, task := range tasks {
		if _, ok := freq[task]; ok {
			freq[task]++
		} else {
			freq[task] = 1
		}
	}

	for key, val := range freq {
		heap.Push(h, Item{
			Value: key,
			Priority: val,
		})
	}

	var time int
	for h.Len() > 0 || len(q) > 0 {
		time++

		for len(q) > 0 && q[0].CooldDownEnd == time {
			v := q[0]
			heap.Push(h, v.HeapItem)
			q = q[1:]
		}

		if h.Len() > 0 {
			item := heap.Pop(h).(Item)
			out = append(out, item.Value)
			if (item.Priority - 1) > 0 {
				q = append(q, QueueItem{
					HeapItem: Item {
						Value: item.Value,
						Priority: item.Priority - 1,
					},
					CooldDownEnd: time + n + 1,
				})
			}
		}
	}

	return time
}
