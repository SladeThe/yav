package internal

import (
	"github.com/SladeThe/yav"
)

func RegisterValidateFunc[K int | string, V any](
	validateFuncs *map[K]yav.ValidateFunc[V],
	key K,
	validateFunc yav.ValidateFunc[V],
) yav.ValidateFunc[V] {
	src := *validateFuncs

	if _, ok := src[key]; ok {
		return validateFunc
	}

	dst := make(map[K]yav.ValidateFunc[V], len(src)+1)

	for k, v := range src {
		dst[k] = v
	}

	dst[key] = validateFunc
	*validateFuncs = dst
	return validateFunc
}

func IsValid[T any](string, T) (bool, error) {
	return false, nil
}
