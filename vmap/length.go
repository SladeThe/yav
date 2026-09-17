package vmap

import (
	"strconv"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

type betweenKey struct {
	min, max int
}

var (
	minFuncs     map[key[int]]any
	maxFuncs     map[key[int]]any
	betweenFuncs map[key[betweenKey]]any
)

func Min[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	k := newKey[int, M](parameter)

	if validateFunc, ok := minFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[int], any](&minFuncs, k, minimum[M](parameter)).(yav.ValidateFunc[M])
}

func Max[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	k := newKey[int, M](parameter)

	if validateFunc, ok := maxFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[int], any](&maxFuncs, k, maximum[M](parameter)).(yav.ValidateFunc[M])
}

func Between[M ~map[K]V, K comparable, V any](minimum, maximum int) yav.ValidateFunc[M] {
	k := newKey[betweenKey, M](betweenKey{min: minimum, max: maximum})

	if validateFunc, ok := betweenFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[betweenKey], any](
		&betweenFuncs, k, between[M](minimum, maximum),
	).(yav.ValidateFunc[M])
}

func minimum[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) < parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func maximum[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) > parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func between[M ~map[K]V, K comparable, V any](minimum, maximum int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) < minimum {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(minimum),
				ValueName: name,
				Value:     value,
			}
		}

		if len(value) > maximum {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(maximum),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
