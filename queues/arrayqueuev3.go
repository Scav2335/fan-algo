package queues

type ArrayQueue[T any] struct {
	items      []T
	head, tail int
}

func NewArrayQueue[T any]() Queue[T] {
	return &ArrayQueue[T]{
		items: make([]T, 1),
	}
}

func (a *ArrayQueue[T]) IsEmpty() bool {
	return a.head == a.tail
}

func (a *ArrayQueue[T]) Enqueue(val T) {
	if a.tail == len(a.items) {
		a.resize(len(a.items) << 1)
	}
	a.items[a.tail] = val
	a.tail++
}

func (a *ArrayQueue[T]) Dequeue() T {
	if a.IsEmpty() {
		panic("Dequeue called on empty queue")
	}

	val := a.items[a.head]
	a.head++

	if !a.IsEmpty() && a.tail-a.head <= len(a.items)/4 {
		a.resize(len(a.items) >> 1)
	}
	return val
}

func (a *ArrayQueue[T]) resize(l int) {
	var (
		ni     = make([]T, l)
		oh, ot = a.head, a.tail
	)

	for i := oh; i < ot; i++ {
		ni[i-oh] = a.items[i]
	}
	a.items = ni
	a.head = 0
	a.tail -= oh
}
