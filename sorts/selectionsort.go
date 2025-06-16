package sorts

func SelectionSort[T any](d []Comparable[T]) {
	l := len(d)
	for i := 0; i < l; i++ {
		mini := i
		for j := i + 1; j < l; j++ {
			if d[j].LessThan(d[mini].Get()) {
				mini = j
			}
		}

		d[i], d[mini] = d[mini], d[i]
	}
}
