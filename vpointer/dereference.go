package vpointer

import (
	"github.com/SladeThe/yav"
)

func Dereference[T any](validateFuncs ...yav.ValidateFunc[T]) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		var yavErrs yav.Errors

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
