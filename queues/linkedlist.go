package queues

type Node struct {
	val  string
	next *Node
}

type LinkedQueueOfStrings struct {
	first *Node
	last  *Node
}

func (l *LinkedQueueOfStrings) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedQueueOfStrings) Enqueue(val string) {
	n := &Node{val: val}
	if l.IsEmpty() {
		l.first = n
	}

	if l.last != nil {
		l.last.next = n
	}
	l.last = n
}

func (l *LinkedQueueOfStrings) Dequeue() string {
	if l.IsEmpty() {
		panic("queue is empty")
	}

	first := l.first
	l.first = first.next
	return first.val
}
