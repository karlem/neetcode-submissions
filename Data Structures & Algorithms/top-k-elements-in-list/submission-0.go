type pair struct {
	val int
	freq int
}

type MinHeap []pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(a, b int) bool {
	return h[a].freq < h[b].freq
}

func (h MinHeap) Swap(a, b int){
	h[a], h[b] = h[b], h[a]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	l := len(old)
	pop := old[l-1]
	old = old[:l-1]
	*h = old
	return pop
}

func topKFrequent(nums []int, k int) []int {
	freqs := map[int]int{}
	for _, num := range nums {
		freqs[num]++
	}

	h := &MinHeap{}

	for val, freq := range freqs {
		heap.Push(h, pair{val: val, freq: freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var out []int
	for _, pair := range *h {
		out = append(out, pair.val)
	}

	return out
}
