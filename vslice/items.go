package vslice

import (
	"fmt"
	"strings"

	"github.com/SladeThe/yav"
)

func Items[S ~[]T, T any](validateFuncs ...yav.ValidateFunc[T]) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		var yavErrs yav.Errors

		for i, v := range value {
			itemChain(&yavErrs, name, i, v, validateFuncs)
		}

		return false, yavErrs.AsError()
	}
}

func itemChain[T any](yavErrs *yav.Errors, name string, index int, value T, validateFuncs []yav.ValidateFunc[T]) {
	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, value)
		yavErrs.AppendConverted(err, withIndex(name, index, err))
		if stop {
			break
		}
	}
}

func withIndex(name string, index int, err error) func(yav.Error) yav.Error {
	if err == nil {
		return nil
	}

	switch typedErr := err.(type) {
	case yav.Error:
	case yav.Errors:
		if len(typedErr.Validation) == 0 {
			return nil
		}
	default:
		return nil
	}

	return withIndexYAV(name, index)
}

func withIndexYAV(name string, index int) func(yav.Error) yav.Error {
	return func(yavErr yav.Error) yav.Error {
		if !strings.HasPrefix(yavErr.ValueName, name) {
			return yavErr
		}

		suffix := yavErr.ValueName[len(name):]

		if suffix != "" && suffix[0] != '.' && suffix[0] != '[' {
			if name != "" {
				return yavErr
			}

			yavErr.ValueName = fmt.Sprintf("[%d].%s", index, suffix)
		} else {
			yavErr.ValueName = fmt.Sprintf("%s[%d]%s", name, index, suffix)
		}

		return yavErr
	}
}
