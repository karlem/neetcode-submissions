type Pair struct {
    time int
    node int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make(map[int][][2]int)
    for _, t := range times {
        u, v, w := t[0], t[1], t[2]
        graph[u] = append(graph[u], [2]int{v, w})
    }

    visited := make(map[int]bool)
    h := &MinHeap{}
    heap.Init(h)
    heap.Push(h, Pair{time: 0, node: k})

    answer := 0

    for h.Len() > 0 {
        curr := heap.Pop(h).(Pair)

        if visited[curr.node] {
            continue
        }

        visited[curr.node] = true
        answer = curr.time

        for _, edge := range graph[curr.node] {
            nextNode, w := edge[0], edge[1]
            if !visited[nextNode] {
                heap.Push(h, Pair{time: curr.time + w, node: nextNode})
            }
        }
    }

    if len(visited) != n {
        return -1
    }
    return answer
}