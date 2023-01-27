package yav

import (
	"go.uber.org/multierr"
)

type Validatable interface {
	Validate() error
}

type Zeroer interface {
	IsZero() bool
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
	if len(validateFuncs) == 1 {
		return validateFuncs[0]
	}

	return func(name string, value T) (stop bool, err error) {
		for _, validateFunc := range validateFuncs {
			if stop, err = validateFunc(name, value); err == nil {
				return
			}
		}

		return
	}
}

func Or2[T any](validateFunc1, validateFunc2 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc1(name, value); err == nil {
			return
		}

		return validateFunc2(name, value)
	}
}

func Or3[T any](validateFunc1, validateFunc2, validateFunc3 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc1(name, value); err == nil {
			return
		}

		if stop, err = validateFunc2(name, value); err == nil {
			return
		}

		return validateFunc3(name, value)
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

		if validationErr.ValueName[0] == '[' {
			validationErr.ValueName = name + validationErr.ValueName
		} else {
			validationErr.ValueName = name + "." + validationErr.ValueName
		}

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

func UnnamedValidate[T Validatable](_ string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, err
}

func NestedValidate[T Validatable](name string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, Nested(name, err)
}
