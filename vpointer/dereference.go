package vpointer

import (
	"github.com/SladeThe/yav"
)

// Dereference applies the given validation funcs to the value a pointer points to.
// It requires non-nil pointer, otherwise fails with yav.CheckNameRequired.
// If pointer is optional, you may precede Dereference with OmitEmpty.
func Dereference[T any](validateFuncs ...yav.ValidateFunc[T]) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		var yavErrs yav.Errors

		if value == nil {
			yavErrs.Append(yav.ErrRequired(name))
			return true, yavErrs.AsError()
		}

		for _, validateFunc := range validateFuncs {
			stop, err = validateFunc(name, *value)
			yavErrs.Append(err)
			if stop {
				break
			}
		}

		return stop, yavErrs.AsError()
	}
}
