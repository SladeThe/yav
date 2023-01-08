package vslice

import (
	"strconv"

	"github.com/SladeThe/yav"
)

func Min[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) < parameter {
			return false, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func Max[S ~[]T, T any](parameter int) yav.ValidateFunc[S] {
	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) > parameter {
			return false, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
