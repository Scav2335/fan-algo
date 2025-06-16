package sorts

const cutoff = 4

func MergeSort[T any](d []Comparable[T]) {
	mergeSort(d, make([]Comparable[T], len(d)), 0, len(d)-1)
}

func BottomUpMergeSort[T any](d []Comparable[T]) {
	l := len(d)
	aux := make([]Comparable[T], l)
	for sz := 1; sz < l; sz = sz << 1 {
		for lo := 0; lo+sz < l; lo += 2 * sz {
			hi := lo + 2*sz - 1
			if hi > l-1 {
				hi = l - 1
			}
			merge(d, aux, lo, lo+sz-1, hi)
		}
	}
}

func mergeSort[T any](a, aux []Comparable[T], lo, hi int) {
	if lo+cutoff-1 >= hi {
		InsertionSortPart(a, lo, hi)
		return
	}
	mid := (lo + hi) / 2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	if !a[mid+1].LessThan(a[mid].Get()) {
		return
	}

	merge(a, aux, lo, mid, hi)
}

func merge[T any](a, aux []Comparable[T], lo, mid, hi int) {
	for i := lo; i < hi+1; i++ {
		aux[i] = a[i]
	}

	for i, j, k := lo, mid+1, lo; k < hi+1; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j].LessThan(aux[i].Get()) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
