package vslice

import (
	"github.com/SladeThe/yav/internal"
)

func Unique[S ~[]T, T comparable](name string, value S) (stop bool, err error) {
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

func isNotUnique2[S ~[]T, T comparable](value S) bool {
	return value[0] == value[1]
}

func isNotUnique3[S ~[]T, T comparable](value S) bool {
	return value[0] == value[1] || value[0] == value[2] || value[1] == value[2]
}

func isNotUnique4[S ~[]T, T comparable](value S) bool {
	return value[0] == value[1] || value[0] == value[2] || value[0] == value[3] ||
		value[1] == value[2] || value[1] == value[3] || value[2] == value[3]
}

func isNotUnique5[S ~[]T, T comparable](value S) bool {
	return value[0] == value[1] || value[0] == value[2] || value[0] == value[3] || value[0] == value[4] ||
		value[1] == value[2] || value[1] == value[3] || value[1] == value[4] ||
		value[2] == value[3] || value[2] == value[4] || value[3] == value[4]
}

func isNotUniqueN[S ~[]T, T comparable](value S) bool {
	m := make(map[T]struct{}, len(value)-1)

	for i, item := range value[:len(value)-1] {
		if m[item] = struct{}{}; len(m) != i+1 {
			return true
		}
	}

	_, ok := m[value[len(value)-1]]
	return ok
}
