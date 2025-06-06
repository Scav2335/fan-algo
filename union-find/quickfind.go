package union_find

type QuickFindUF struct {
	ids []int
}

func NewQuickFindUF(size int) *QuickFindUF {
	ids := make([]int, size)
	for i := 0; i < size; i++ {
		ids[i] = i
	}
	return &QuickFindUF{ids: ids}
}

func (u *QuickFindUF) Connected(p, q int) bool {
	return u.ids[p] == u.ids[q]
}

func (u *QuickFindUF) Union(p, q int) {
	pid, qid := u.ids[p], u.ids[q]
	for i, id := range u.ids {
		if id == pid {
			u.ids[i] = qid
		}
	}
}
