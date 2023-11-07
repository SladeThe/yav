package vmap

import (
	"github.com/SladeThe/yav/internal"
)

func Unique[M ~map[K]V, K comparable, V comparable](name string, value M) (stop bool, err error) {
	switch len(value) {
	case 0, 1:
	case 2:
		if isNotUnique2(value) {
			return true, internal.ErrUnique(name, value)
		}
	case 3:
		if isNotUnique3(value) {
			return true, internal.ErrUnique(name, value)
		}
	case 4:
		if isNotUnique4(value) {
			return true, internal.ErrUnique(name, value)
		}
	case 5:
		if isNotUnique5(value) {
			return true, internal.ErrUnique(name, value)
		}
	default:
		if isNotUniqueN(value) {
			return true, internal.ErrUnique(name, value)
		}
	}

	return false, nil
}

func isNotUnique2[M ~map[K]V, K comparable, V comparable](value M) bool {
	var values [1]V
	i := 0

	for _, v := range value {
		if i > 0 {
			return v == values[0]
		}

		values[0] = v
		i++
	}

	panic("unreachable")
}

func isNotUnique3[M ~map[K]V, K comparable, V comparable](value M) bool {
	var values [2]V
	i := 0

	for _, v := range value {
		switch i {
		case 2:
			return v == values[0] || v == values[1]
		case 1:
			if v == values[0] {
				return true
			}
		}

		values[i] = v
		i++
	}

	panic("unreachable")
}

func isNotUnique4[M ~map[K]V, K comparable, V comparable](value M) bool {
	var values [3]V
	i := 0

	for _, v := range value {
		switch i {
		case 3:
			return v == values[0] || v == values[1] || v == values[2]
		case 2:
			if v == values[0] || v == values[1] {
				return true
			}
		case 1:
			if v == values[0] {
				return true
			}
		}

		values[i] = v
		i++
	}

	panic("unreachable")
}

func isNotUnique5[M ~map[K]V, K comparable, V comparable](value M) bool {
	var values [4]V
	i := 0

	for _, v := range value {
		switch i {
		case 4:
			return v == values[0] || v == values[1] || v == values[2] || v == values[3]
		case 3:
			if v == values[0] || v == values[1] || v == values[2] {
				return true
			}
		case 2:
			if v == values[0] || v == values[1] {
				return true
			}
		case 1:
			if v == values[0] {
				return true
			}
		}

		values[i] = v
		i++
	}

	panic("unreachable")
}

func isNotUniqueN[M ~map[K]V, K comparable, V comparable](value M) bool {
	m := make(map[V]struct{}, len(value)-1)
	i := 0

	for _, v := range value {
		if i == len(value)-1 {
			_, ok := m[v]
			return ok
		}

		m[v] = struct{}{}
		i++

		if len(m) != i {
			return true
		}
	}

	panic("unreachable")
}
