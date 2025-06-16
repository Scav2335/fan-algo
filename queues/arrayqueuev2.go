package queues

type ArrayQueueOfStringsV2 struct {
	items      []string
	head, tail int
}

func (a *ArrayQueueOfStringsV2) IsEmpty() bool {
	return a.head == a.tail
}

func (a *ArrayQueueOfStringsV2) Enqueue(val string) {
	if a.tail == len(a.items) {
		a.resize(len(a.items) << 1)
	}
	a.items[a.tail] = val
	a.tail++
}

func (a *ArrayQueueOfStringsV2) Dequeue() string {
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

func (a *ArrayQueueOfStringsV2) resize(l int) {
	var (
		ni     = make([]string, l)
		oh, ot = a.head, a.tail
	)

	for i := oh; i < ot; i++ {
		ni[i-oh] = a.items[i]
	}
	a.items = ni
	a.head = 0
	a.tail -= oh
}
