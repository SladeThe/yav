package vmap

import (
	"github.com/SladeThe/yav/internal"
)

func Unique[M ~map[K]V, K comparable, V comparable](name string, value M) (stop bool, err error) {
	if len(value) <= 1 {
		return false, nil
	}

	m := make(map[V]struct{}, len(value)-1)

	expectedLength := 0

	for _, v := range value {
		if expectedLength == len(value)-1 {
			if _, ok := m[v]; ok {
				return true, internal.ErrUnique(name, value)
			}

			return false, nil
		}

		m[v] = struct{}{}
		expectedLength++

		if len(m) != expectedLength {
			return true, internal.ErrUnique(name, value)
		}
	}

	return false, nil
}
