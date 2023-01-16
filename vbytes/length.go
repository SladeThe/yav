package vbytes

import (
	"strconv"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

type betweenKey struct {
	min, max int
}

var (
	minFuncs     map[int]yav.ValidateFunc[[]byte]
	maxFuncs     map[int]yav.ValidateFunc[[]byte]
	betweenFuncs map[betweenKey]yav.ValidateFunc[[]byte]
)

func Min(parameter int) yav.ValidateFunc[[]byte] {
	if validateFunc, ok := minFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&minFuncs, parameter, min(parameter))
}

func Max(parameter int) yav.ValidateFunc[[]byte] {
	if validateFunc, ok := maxFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&maxFuncs, parameter, max(parameter))
}

func Between(min, max int) yav.ValidateFunc[[]byte] {
	k := betweenKey{min: min, max: max}

	if validateFunc, ok := betweenFuncs[k]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&betweenFuncs, k, between(min, max))
}

func min(parameter int) yav.ValidateFunc[[]byte] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) < parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: parameterString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func max(parameter int) yav.ValidateFunc[[]byte] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: parameterString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func between(min, max int) yav.ValidateFunc[[]byte] {
	minString := strconv.Itoa(min)
	maxString := strconv.Itoa(max)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) < min {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: minString,
				ValueName: name,
				Value:     value,
			}
		}

		if len(value) > max {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: maxString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
