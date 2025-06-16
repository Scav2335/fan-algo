package stacks

type Stack[T any] interface {
	IsEmpty() bool
	Push(val T)
	Pop() T
}
