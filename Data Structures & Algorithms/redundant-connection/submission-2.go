type UnionFind struct {
	parents map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parents: map[int]int{},
	}
}

func (u *UnionFind) Find(x int) int {
	if _, ok := u.parents[x]; !ok {
		u.parents[x] = x
	}

	if u.parents[x] == x {
		return x
	}

	u.parents[x] = u.Find(u.parents[x])
	return u.parents[x]
}

func (u *UnionFind) Union(a, b int) {
	rootA := u.Find(a)
	rootB := u.Find(b)

	if rootA != rootB {
		u.parents[rootB] = rootA
	}
}

func (u *UnionFind) Connected(a, b int) bool {
	return u.Find(a) == u.Find(b)
}

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind()

	var res []int
    for _, edge := range edges {
		a, b := edge[0], edge[1]

		if uf.Connected(a, b) {
			res = []int{a,b}
			continue
		}

		uf.Union(a, b)
	}

	return res
}
