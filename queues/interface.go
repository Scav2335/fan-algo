package queues

type Queue[T any] interface {
	IsEmpty() bool
	Enqueue(val T)
	Dequeue() T
}
