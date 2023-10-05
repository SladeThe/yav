package vpointer

import (
	"github.com/SladeThe/yav"
)

func Dereference[T any](validateFuncs ...yav.ValidateFunc[T]) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		var yavErrs yav.Errors

		if value == nil {
			yavErrs.Append(yav.Error{
				CheckName: yav.CheckNameRequired,
				ValueName: name,
			})

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
