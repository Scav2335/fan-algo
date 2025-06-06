package union_find

type QuickUnionUF struct {
	ids []int
	sz  []int
}

func NewQuickUnionUF(size int) *QuickUnionUF {
	ids, sz := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		ids[i] = i
		sz[i] = 1
	}
	return &QuickUnionUF{
		ids: ids,
		sz:  sz,
	}
}

// weighted quick-union with path compression
// depth  of any node x is at most lgN
func (u *QuickUnionUF) root(i int) int {
	for i != u.ids[i] {
		u.ids[i] = u.ids[u.ids[i]]
		i = u.ids[i]
	}
	return i
}

func (u *QuickUnionUF) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *QuickUnionUF) Union(p, q int) {
	i, j := u.root(p), u.root(q)
	if i == j {
		return
	}

	if u.sz[i] < u.sz[j] {
		u.ids[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.ids[j] = i
		u.sz[i] += u.sz[j]
	}
}
