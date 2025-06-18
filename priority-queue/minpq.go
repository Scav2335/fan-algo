package priority_queue

import "scav.abc/fan-algo/sorts"

type MinPQ[T sorts.Comparable[T]] struct {
	items []*T
	N     int
}

func NewMinPQ[T sorts.Comparable[T]](m int) PriorityQueue[T] {
	return &MinPQ[T]{
		items: make([]*T, m+1),
	}
}

func (m *MinPQ[T]) Insert(a T) {
	m.N++
	m.items[m.N] = &a
	m.swim(m.N)
}

func (m *MinPQ[T]) Delete() T {
	r := m.items[1]
	m.items[1], m.items[m.N] = m.items[m.N], m.items[1]
	m.N--
	m.sink(1)

	return *r
}

func (m *MinPQ[T]) Size() int {
	return m.N
}

func (m *MinPQ[T]) IsEmpty() bool {
	return m.N == 0
}

func (m *MinPQ[T]) swim(k int) {
	for k > 1 {
		p := k >> 1
		if p < 1 || !(*m.items[p]).GreaterThan(*m.items[k]) {
			break
		}
		m.items[p], m.items[k] = m.items[k], m.items[p]
		k = p
	}
}

func (m *MinPQ[T]) sink(k int) {
	for k < m.N {
		c := k << 1
		if c < m.N && (*m.items[c+1]).LessThan(*m.items[c]) {
			c++
		}
		if c > m.N || !(*m.items[k]).GreaterThan(*m.items[c]) {
			break
		}
		m.items[k], m.items[c] = m.items[c], m.items[k]
		k = c
	}
}

func (m *MinPQ[T]) Range(fn func(i int, a T) bool) {
	for k := 1; k <= m.N; k++ {
		if m.items[k] == nil {
			continue
		}
		ok := fn(k, *m.items[k])
		if !ok {
			break
		}
	}
}
