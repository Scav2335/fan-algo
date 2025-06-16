package queues

type ArrayQueueOfStrings struct {
	items []string
}

func (a *ArrayQueueOfStrings) IsEmpty() bool {
	return len(a.items) == 0
}

func (a *ArrayQueueOfStrings) Enqueue(val string) {
	a.items = append(a.items, val)
}

func (a *ArrayQueueOfStrings) Dequeue() string {
	if a.IsEmpty() {
		panic("queue is empty")
	}
	val := a.items[0]
	a.items = a.items[1:]

	return val
}
