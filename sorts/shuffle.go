package sorts

import "math/rand/v2"

func Shuffle[T comparable](s []T) {
	l := len(s)
	for i := 0; i < l; i++ {
		r := rand.IntN(i + 1)
		s[i], s[r] = s[r], s[i]
	}
}
