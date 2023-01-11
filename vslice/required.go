package vslice

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

func OmitEmpty[S ~[]T, T any](_ string, value S) (stop bool, err error) {
	return len(value) == 0, nil
}

func Required[S ~[]T, T any](name string, value S) (stop bool, err error) {
	if len(value) == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredWithAny[S ~[]T, T any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[S] {
	if !accumulator.IsEnabled() {
		return internal.Valid[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithoutAny[S ~[]T, T any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[S] {
	if !accumulator.IsEnabled() {
		return internal.Valid[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithAll[S ~[]T, T any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[S] {
	if !accumulator.IsEnabled() {
		return internal.Valid[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithoutAll[S ~[]T, T any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[S] {
	if !accumulator.IsEnabled() {
		return internal.Valid[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}
