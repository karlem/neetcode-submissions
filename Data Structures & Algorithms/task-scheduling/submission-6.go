type Item struct {
    char byte
    cnt  int
}

type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
    return h[i].cnt > h[j].cnt
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type Cooldown struct {
    item  Item
    time  int
}

func leastInterval(tasks []byte, n int) int {
    freq := make(map[byte]int)
    for _, t := range tasks {
        freq[t]++
    }

    h := &MaxHeap{}
    for ch, c := range freq {
        heap.Push(h, Item{char: ch, cnt: c})
    }

    cooldown := []Cooldown{}
    time := 0

    for h.Len() > 0 || len(cooldown) > 0 {
        time++

        // release tasks whose cooldown is done
        for len(cooldown) > 0 && cooldown[0].time == time {
            heap.Push(h, cooldown[0].item)
            cooldown = cooldown[1:]
        }

        // execute a task if available
        if h.Len() > 0 {
            cur := heap.Pop(h).(Item)
            cur.cnt--

            if cur.cnt > 0 {
                cooldown = append(cooldown, Cooldown{
                    item: cur,
                    time: time + n + 1,
                })
            }
        }
    }

    return time
}