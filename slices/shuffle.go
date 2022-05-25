package slices

import (
	"math/rand"
	"time"
)

func Shuffle[T any](slice []T) []T {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	size := len(slice)
	perm := r.Perm(size)
	ret := make([]T, size)

	for i, r := range perm {
		ret[i] = slice[r]
	}

	return ret
}
