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

func QuickSortThreePart[T any](d []Comparable[T]) {
	Shuffle(d)
	sortThreePart(d, 0, len(d)-1)
}

func sortThreePart[T any](d []Comparable[T], lo, hi int) {
	if lo >= hi {
		return
	}
	lt, gt, i, v := lo, hi, lo, d[lo]
	for i <= gt {
		cmp := d[i].CompareTo(v.Get())
		if cmp < 0 {
			d[i], d[lt] = d[lt], d[i]
			i++
			lt++
		} else if cmp > 0 {
			d[i], d[gt] = d[gt], d[i]
			gt--
		} else {
			i++
		}
	}
	sortThreePart(d, lo, lt-1)
	sortThreePart(d, gt+1, hi)
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

func Select[T any](d []Comparable[T], k int) Comparable[T] {
	Shuffle(d)
	lo, hi := 0, len(d)-1

	for hi > lo {
		p := partition(d, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return d[k]
		}
	}

	return d[k]
}
