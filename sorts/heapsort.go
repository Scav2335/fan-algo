package sorts

func HeapSort[T any](d []Comparable[T]) {
	N := len(d) - 1
	// bottom to up
	for i := N >> 1; i >= 1; i-- {
		sink(d, i, N)
	}

	for N > 1 {
		d[1], d[N] = d[N], d[1]
		N--
		sink(d, 1, N)
	}
}

func sink[T any](a []Comparable[T], k, N int) {
	for k < N {
		j := k << 1
		if j < N && a[j].LessThan(a[j+1].Get()) {
			j++
		}
		if j > N || !a[k].LessThan(a[j].Get()) {
			break
		}

		a[j], a[k] = a[k], a[j]
		k = j
	}
}
