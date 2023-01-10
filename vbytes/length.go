package vbytes

import (
	"strconv"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

var (
	minFuncs map[int]yav.ValidateFunc[[]byte]
	maxFuncs map[int]yav.ValidateFunc[[]byte]
)

func Min(parameter int) yav.ValidateFunc[[]byte] {
	if validateFunc, ok := minFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&minFuncs, parameter, min(parameter))
}

func min(parameter int) yav.ValidateFunc[[]byte] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) < parameter {
			return false, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: parameterString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func Max(parameter int) yav.ValidateFunc[[]byte] {
	if validateFunc, ok := maxFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&maxFuncs, parameter, max(parameter))
}

func max(parameter int) yav.ValidateFunc[[]byte] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > parameter {
			return false, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: parameterString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
