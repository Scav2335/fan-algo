package sorts

type Comparable[T any] interface {
	Get() T
	LessThan(a T) bool
}
