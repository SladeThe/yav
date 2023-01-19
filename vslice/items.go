package vslice

import (
	"fmt"
	"strings"

	"go.uber.org/multierr"

	"github.com/SladeThe/yav"
)

func Items[S ~[]T, T any](validateFuncs ...yav.ValidateFunc[T]) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		for i, v := range value {
			multierr.AppendInto(&err, itemChain(name, i, v, validateFuncs))
		}

		return
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

	if validationErr, ok := err.(yav.Error); ok {
		if !strings.HasPrefix(validationErr.ValueName, name) {
			return err
		}

		validationErr.ValueName = fmt.Sprintf("%s[%d]%s", name, index, validationErr.ValueName[len(name):])
		return validationErr
	}

	partialErrs := multierr.Errors(err)
	if len(partialErrs) <= 1 {
		return err
	}

	var combinedErr error

	for _, partialErr := range partialErrs {
		multierr.AppendInto(&combinedErr, withIndex(name, index, partialErr))
	}

	return combinedErr
}
