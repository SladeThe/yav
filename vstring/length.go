package vstring

import (
	"strconv"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

var (
	minFuncs = make(map[int]yav.ValidateFunc[string])
	maxFuncs = make(map[int]yav.ValidateFunc[string])
)

func Min(parameter int) yav.ValidateFunc[string] {
	if validateFunc, ok := minFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterValidateFunc(&minFuncs, parameter, min(parameter))
}

func min(parameter int) yav.ValidateFunc[string] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value string) (stop bool, err error) {
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

func Max(parameter int) yav.ValidateFunc[string] {
	if validateFunc, ok := maxFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterValidateFunc(&maxFuncs, parameter, max(parameter))
}

func max(parameter int) yav.ValidateFunc[string] {
	parameterString := strconv.Itoa(parameter)

	return func(name string, value string) (stop bool, err error) {
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
