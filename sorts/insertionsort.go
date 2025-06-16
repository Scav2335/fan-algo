package sorts

type Int int

func (one Int) LessThan(that Int) bool {
	return one < that
}

func (one Int) Get() Int {
	return one
}

type String string

func (one String) LessThan(that String) bool {
	return one < that
}

func (one String) Get() String {
	return one
}

func InsertionSort[T any](d []Comparable[T]) {
	l := len(d)
	for i := 0; i < l; i++ {
		for j := i; j > 0; j-- {
			if d[j].LessThan(d[j-1].Get()) {
				d[j], d[j-1] = d[j-1], d[j]
			} else {
				// break
				break
			}
		}
	}
}

func InsertionSortPart[T any](d []Comparable[T], lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo; j-- {
			if d[j].LessThan(d[j-1].Get()) {
				d[j], d[j-1] = d[j-1], d[j]
			} else {
				// break
				break
			}
		}
	}
}
