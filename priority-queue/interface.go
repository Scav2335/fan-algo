package priority_queue

import "scav.abc/fan-algo/sorts"

type PriorityQueue[T sorts.Comparable[T]] interface {
	Insert(a T)
	Delete() T
	Size() int
	IsEmpty() bool
	Range(fn func(i int, a T) bool)
}
