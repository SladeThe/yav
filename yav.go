package yav

import (
	"go.uber.org/multierr"
)

type Validatable interface {
	Validate() error
}

type ValidateFunc[T any] func(name string, value T) (stop bool, err error)

func Chain[T any](name string, value T, validateFuncs ...ValidateFunc[T]) error {
	var combinedErr error

	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, value)
		if err != nil {
			multierr.AppendInto(&combinedErr, err)
		}
		if stop {
			break
		}
	}

	return combinedErr
}

func Or[T any](validateFuncs ...ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		for _, validateFunc := range validateFuncs {
			if stop, err = validateFunc(name, value); err == nil {
				return
			}
		}

		return
	}
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

	partialErrs := multierr.Errors(err)
	if len(partialErrs) <= 1 {
		return err
	}

	var combinedErr error

	for _, partialErr := range partialErrs {
		multierr.AppendInto(&combinedErr, Nested(name, partialErr))
	}

	return combinedErr
}

func CallValidate[T Validatable](_ string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, err
}

func NestedCallValidate[T Validatable](name string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, Nested(name, err)
}
