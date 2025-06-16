package stacks

type ArrayStack[T any] struct {
	items []T
	n     int
}

func NewArrayStack[T any]() Stack[T] {
	return &ArrayStack[T]{
		items: make([]T, 1),
	}
}

func (a *ArrayStack[T]) IsEmpty() bool {
	return a.n == 0
}

func (a *ArrayStack[T]) Push(val T) {
	if a.n == len(a.items) {
		a.resize(len(a.items) << 1)
	}

	a.items[a.n] = val
	a.n++
}

func (a *ArrayStack[T]) Pop() T {
	if a.IsEmpty() {
		panic("stack is empty")
	}

	a.n--
	val := a.items[a.n]

	if !a.IsEmpty() && a.n <= len(a.items)/4 {
		a.resize(len(a.items) >> 1)
	}
	return val
}

func (a *ArrayStack[T]) resize(l int) {
	ni := make([]T, l)
	for i := 0; i < a.n; i++ {
		ni[i] = a.items[i]
	}
	a.items = ni
}
