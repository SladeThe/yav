package vslice

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

func Min[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	k := newKey[int, S](parameter)

	if validateFunc, ok := minFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[int], any](&minFuncs, k, min[S](parameter)).(yav.ValidateFunc[S])
}

func Max[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	k := newKey[int, S](parameter)

	if validateFunc, ok := maxFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[int], any](&maxFuncs, k, max[S](parameter)).(yav.ValidateFunc[S])
}

func Between[S ~[]T, T any](min, max int) yav.ValidateFunc[S] {
	k := newKey[betweenKey, S](betweenKey{min: min, max: max})

	if validateFunc, ok := betweenFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[betweenKey], any](&betweenFuncs, k, between[S](min, max)).(yav.ValidateFunc[S])
}

func min[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
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

func max[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
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

func between[S ~[]T, T any](min, max int) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		if len(value) < min {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(min),
				ValueName: name,
				Value:     value,
			}
		}

		if len(value) > max {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(max),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
