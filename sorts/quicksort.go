package sorts

func QuickSort[T any](d []Comparable[T]) {
	Shuffle(d)
	sort(d, 0, len(d)-1)
}

func sort[T any](d []Comparable[T], lo, hi int) {
	if lo+cutoff-1 >= hi {
		InsertionSortPart(d, lo, hi)
		return
	}
	p := partition(d, lo, hi)
	sort(d, lo, p-1)
	sort(d, p+1, hi)
}

func partition[T any](a []Comparable[T], lo, hi int) int {
	i, j := lo, hi+1

	for {
		for {
			i++
			if i == hi || !a[i].LessThan(a[lo].Get()) {
				break
			}
		}
		for {
			j--
			if j == lo || !a[lo].LessThan(a[j].Get()) {
				break
			}
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]

	return j
}
