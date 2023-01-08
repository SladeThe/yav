package yav

import (
	"go.uber.org/multierr"
)

type ValidateFunc[T any] func(name string, value T) (stop bool, err error)

func Chain[T any](name string, value T, validateFuncs ...ValidateFunc[T]) error {
	for _, validateFunc := range validateFuncs {
		if stop, err := validateFunc(name, value); stop || err != nil {
			return err
		}
	}

	return nil
}

func Or[T any](name string, value T, validateFuncs ...ValidateFunc[T]) (stop bool, err error) {
	for _, validateFunc := range validateFuncs {
		if stop, err = validateFunc(name, value); stop || err == nil {
			return
		}
	}

	return
}

func Nested(name string, err error) error {
	if err == nil {
		return nil
	}

	if validationErr, ok := err.(Error); ok {
		if validationErr.ValueName == "" {
			return err
		}

		validationErr.ValueName = name + "." + validationErr.ValueName
		return validationErr
	}

	partialErrs := multierr.Errors(err) // TODO update in-place ?
	if len(partialErrs) <= 1 {
		return err
	}

	var combinedErr error

	for _, partialErr := range partialErrs {
		multierr.AppendInto(&combinedErr, Nested(name, partialErr))
	}

	return combinedErr
}
