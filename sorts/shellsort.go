package sorts

var sorting = []int{7, 3, 1}

func ShellSort[T any](d []Comparable[T]) {
	for _, h := range sorting {
		insertionSort(d, h)
	}
}

func insertionSort[T any](d []Comparable[T], h int) {
	l := len(d)
	for i := h; i < l; i++ {
		for j := i; j >= h; j -= h {
			if d[j].LessThan(d[j-h].Get()) {
				d[j], d[j-h] = d[j-h], d[j]
			} else {
				// break
				break
			}
		}
	}
}
