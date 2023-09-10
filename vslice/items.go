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
			yavErrs.Append(itemChain(name, i, v, validateFuncs))
		}

		return false, yavErrs.AsError()
	}
}

func itemChain[T any](name string, index int, value T, validateFuncs []yav.ValidateFunc[T]) error {
	for _, validateFunc := range validateFuncs {
		if stop, err := validateFunc(name, value); stop {
			return withIndex(name, index, err)
		}
	}

	return nil
}

func withIndex(name string, index int, err error) error {
	if err == nil {
		return nil
	}

	switch typedErr := err.(type) {
	case yav.Error:
		return withIndexYAV(name, index, typedErr)
	case yav.Errors:
		for i, yavErr := range typedErr.Validation {
			typedErr.Validation[i] = withIndexYAV(name, index, yavErr)
		}

		return typedErr
	default:
		return err
	}
}

func withIndexYAV(name string, index int, yavErr yav.Error) yav.Error {
	if !strings.HasPrefix(yavErr.ValueName, name) {
		return yavErr
	}

	if len(yavErr.ValueName) > len(name) && yavErr.ValueName[len(name)] != '.' {
		yavErr.ValueName = fmt.Sprintf("%s[%d].%s", name, index, yavErr.ValueName[len(name):])
	} else {
		yavErr.ValueName = fmt.Sprintf("%s[%d]%s", name, index, yavErr.ValueName[len(name):])
	}

	return yavErr
}
