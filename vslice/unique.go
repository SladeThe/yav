package vslice

import (
	"github.com/SladeThe/yav"
)

func IsUnique[S ~[]T, T comparable](name string, value S) (stop bool, err error) {
	switch len(value) {
	case 0, 1:
	case 2:
		if value[0] == value[1] {
			return false, uniqueErr(name, value)
		}
	case 3:
		if value[0] == value[1] || value[0] == value[2] || value[1] == value[2] {
			return false, uniqueErr(name, value)
		}
	case 4:
		if value[0] == value[1] || value[0] == value[2] || value[0] == value[3] ||
			value[1] == value[2] || value[1] == value[3] || value[2] == value[3] {
			return false, uniqueErr(name, value)
		}
	case 5:
		if value[0] == value[1] || value[0] == value[2] || value[0] == value[3] || value[0] == value[4] ||
			value[1] == value[2] || value[1] == value[3] || value[1] == value[4] ||
			value[2] == value[3] || value[2] == value[4] || value[3] == value[4] {
			return false, uniqueErr(name, value)
		}
	default:
		m := make(map[T]struct{}, len(value))

		for _, item := range value {
			if _, ok := m[item]; ok {
				return false, uniqueErr(name, value)
			}

			m[item] = struct{}{}
		}
	}

	return false, nil
}

func uniqueErr(name string, value any) yav.Error {
	return yav.Error{
		CheckName: yav.CheckNameUnique,
		ValueName: name,
		Value:     value,
	}
}
