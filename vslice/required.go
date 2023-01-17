package vslice

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
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

func RequiredWithAny[S ~[]T, T any](fields string) accumulators.RequiredWithAny[S] {
	return accumulators.NewRequiredWithAny(fields, provideRequiredWithAny[S])
}

func RequiredWithoutAny[S ~[]T, T any](fields string) accumulators.RequiredWithoutAny[S] {
	return accumulators.NewRequiredWithoutAny(fields, provideRequiredWithoutAny[S])
}

func RequiredWithAll[S ~[]T, T any](fields string) accumulators.RequiredWithAll[S] {
	return accumulators.NewRequiredWithAll(fields, provideRequiredWithAll[S])
}

func RequiredWithoutAll[S ~[]T, T any](fields string) accumulators.RequiredWithoutAll[S] {
	return accumulators.NewRequiredWithoutAll(fields, provideRequiredWithoutAll[S])
}

func provideRequiredWithAny[S ~[]T, T any](fields string, enabled bool) yav.ValidateFunc[S] {
	if !enabled {
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

func provideRequiredWithoutAny[S ~[]T, T any](fields string, enabled bool) yav.ValidateFunc[S] {
	if !enabled {
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

func provideRequiredWithAll[S ~[]T, T any](fields string, enabled bool) yav.ValidateFunc[S] {
	if !enabled {
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

func provideRequiredWithoutAll[S ~[]T, T any](fields string, enabled bool) yav.ValidateFunc[S] {
	if !enabled {
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
