package sorts

type Int int

func (one Int) LessThan(that Int) bool {
	return one < that
}

func (one Int) GreaterThan(that Int) bool {
	return one > that
}

func (one Int) CompareTo(that Int) int {
	if one < that {
		return -1
	} else if one > that {
		return 1
	} else {
		return 0
	}
}

func (one Int) Get() Int {
	return one
}

type String string

func (one String) LessThan(that String) bool {
	return one < that
}

func (one String) GreaterThan(that String) bool {
	return one > that
}

func (one String) CompareTo(that String) int {
	if one < that {
		return -1
	} else if one > that {
		return 1
	} else {
		return 0
	}
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
