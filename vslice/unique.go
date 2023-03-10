package vslice

import (
	"github.com/SladeThe/yav/internal"
)

func Unique[S ~[]T, T comparable](name string, value S) (stop bool, err error) {
	switch len(value) {
	case 0, 1:
	case 2:
		if value[0] == value[1] {
			return true, internal.ErrUnique(name, value)
		}
	case 3:
		if value[0] == value[1] || value[0] == value[2] || value[1] == value[2] {
			return true, internal.ErrUnique(name, value)
		}
	case 4:
		if value[0] == value[1] || value[0] == value[2] || value[0] == value[3] ||
			value[1] == value[2] || value[1] == value[3] || value[2] == value[3] {
			return true, internal.ErrUnique(name, value)
		}
	case 5:
		if value[0] == value[1] || value[0] == value[2] || value[0] == value[3] || value[0] == value[4] ||
			value[1] == value[2] || value[1] == value[3] || value[1] == value[4] ||
			value[2] == value[3] || value[2] == value[4] || value[3] == value[4] {
			return true, internal.ErrUnique(name, value)
		}
	default:
		m := make(map[T]struct{}, len(value)-1)

		for i, item := range value[:len(value)-1] {
			if m[item] = struct{}{}; len(m) != i+1 {
				return true, internal.ErrUnique(name, value)
			}
		}

		if _, ok := m[value[len(value)-1]]; ok {
			return true, internal.ErrUnique(name, value)
		}
	}

	return false, nil
}
