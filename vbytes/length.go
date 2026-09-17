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

	return internal.RegisterMapEntry(&minFuncs, parameter, minimum(parameter))
}

func Max(parameter int) yav.ValidateFunc[[]byte] {
	if validateFunc, ok := maxFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&maxFuncs, parameter, maximum(parameter))
}

func Between(minimum, maximum int) yav.ValidateFunc[[]byte] {
	k := betweenKey{min: minimum, max: maximum}

	if validateFunc, ok := betweenFuncs[k]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&betweenFuncs, k, between(minimum, maximum))
}

func minimum(parameter int) yav.ValidateFunc[[]byte] {
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

func maximum(parameter int) yav.ValidateFunc[[]byte] {
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

func between(minimum, maximum int) yav.ValidateFunc[[]byte] {
	minString := strconv.Itoa(minimum)
	maxString := strconv.Itoa(maximum)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) < minimum {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: minString,
				ValueName: name,
				Value:     value,
			}
		}

		if len(value) > maximum {
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
