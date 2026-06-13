type UnionFind struct {
	parents map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parents: map[int]int{},
	}
}

func (uf *UnionFind) Find(i int) int {
	if _, ok := uf.parents[i]; !ok {
		uf.parents[i] = i
		return uf.parents[i]
	}

	if uf.parents[i] != i {
		uf.parents[i] = uf.Find(uf.parents[i])
	}
	return uf.parents[i]
}

func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		uf.parents[rootI] = rootJ
	}
}

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind()
	var res []int
	for _, edge := range edges {
		a, b := edge[0], edge[1]

		rootA := uf.Find(a)
		rootB := uf.Find(b)

		if rootA == rootB {
			res = edge
			continue
		}

		uf.Union(a, b)
	}

	return res
}
