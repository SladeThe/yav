package internal

import (
	"github.com/SladeThe/yav"
)

func RegisterMapEntry[K comparable, V any](m *map[K]V, key K, value V) V {
	src := *m

	if _, ok := src[key]; ok {
		return value
	}

	dst := make(map[K]V, len(src)+1)

	for k, v := range src {
		dst[k] = v
	}

	dst[key] = value
	*m = dst
	return value
}

func Valid[T any](string, T) (bool, error) {
	return false, nil
}

func ErrUnique(name string, value any) yav.Error {
	return yav.Error{
		CheckName: yav.CheckNameUnique,
		ValueName: name,
		Value:     value,
	}
}
