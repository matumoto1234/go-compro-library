package unionfind

import "github.com/matumoto1234/go-compro-library/util/assert"

type UnionFind struct {
	n      int
	parent []int
	size   []int
}

func New(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}

	return &UnionFind{
		n:      n,
		parent: parent,
		size:   size,
	}
}

type MergeState int

const (
	NotMerged MergeState = iota
	LeftMerged
	RightMerged
)

func (u *UnionFind) Merge(l, r int) MergeState {
	assert.Do(0 <= l && l < u.n)
	assert.Do(0 <= r && r < u.n)
	x := u.Root(l)
	y := u.Root(r)
	if x == y {
		return NotMerged
	}

	state := RightMerged
	if u.size[x] < u.size[y] {
		x, y = y, x
		state = LeftMerged
	}

	u.size[x] += u.size[y]
	u.parent[y] = x

	return state
}

func (u *UnionFind) Root(a int) int {
	assert.Do(0 <= a && a < u.n)
	if u.parent[a] == a {
		return a
	}
	u.parent[a] = u.Root(u.parent[a])
	return u.parent[a]
}

func (u *UnionFind) Same(a, b int) bool {
	assert.Do(0 <= a && a < u.n)
	assert.Do(0 <= b && b < u.n)
	return u.Root(a) == u.Root(b)
}

func (u *UnionFind) Size(a int) int {
	assert.Do(0 <= a && a < u.n)
	return u.size[u.Root(a)]
}

// Groups() : Returns the list of connected components.
func (u *UnionFind) Groups() [][]int {
	grp := make(map[int][]int)
	for i := 0; i < u.n; i++ {
		grp[u.Root(i)] = append(grp[u.Root(i)], i)
	}

	res := make([][]int, 0, len(grp))
	for _, v := range grp {
		res = append(res, v)
	}

	return res
}
