package vmap

import (
	"github.com/SladeThe/yav/internal"
)

func Unique[M ~map[K]V, K comparable, V comparable](name string, value M) (stop bool, err error) {
	if len(value) <= 1 {
		return false, nil
	}

	m := make(map[V]struct{}, len(value))

	for _, v := range value {
		if _, ok := m[v]; ok {
			return true, internal.ErrUnique(name, value)
		}

		m[v] = struct{}{} // TODO do not add the last, also reduce map capacity + benchmark
	}

	return false, nil
}
