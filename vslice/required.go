package vslice

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredIfFuncs     map[key[string]]any
	requiredUnlessFuncs map[key[string]]any
)

func OmitEmpty[S ~[]T, T any](_ string, value S) (stop bool, err error) {
	return len(value) == 0, nil
}

func Required[S ~[]T, T any](name string, value S) (stop bool, err error) {
	if len(value) == 0 {
		return true, yav.ErrRequired(name)
	}

	return false, nil
}

func RequiredIf[S ~[]T, T any](conditionString string, condition bool) yav.ValidateFunc[S] {
	if !condition {
		return OmitEmpty[S]
	}

	k := newKey[string, S](conditionString)

	if validateFunc, ok := requiredIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredIfFuncs, k, requiredIf[S](conditionString),
	).(yav.ValidateFunc[S])
}

func RequiredUnless[S ~[]T, T any](conditionString string, condition bool) yav.ValidateFunc[S] {
	if condition {
		return OmitEmpty[S]
	}

	k := newKey[string, S](conditionString)

	if validateFunc, ok := requiredUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredUnlessFuncs, k, requiredUnless[S](conditionString),
	).(yav.ValidateFunc[S])
}

func RequiredWithAny[S ~[]T, T any]() accumulators.RequiredWithAny[S] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAny[S])
}

func RequiredWithoutAny[S ~[]T, T any]() accumulators.RequiredWithoutAny[S] {
	return accumulators.NewRequiredWithoutAny(provideRequiredWithoutAny[S])
}

func RequiredWithAll[S ~[]T, T any]() accumulators.RequiredWithAll[S] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAll[S])
}

func RequiredWithoutAll[S ~[]T, T any]() accumulators.RequiredWithoutAll[S] {
	return accumulators.NewRequiredWithoutAll(provideRequiredWithoutAll[S])
}

func requiredIf[S ~[]T, T any](conditionString string) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredIf,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func requiredUnless[S ~[]T, T any](conditionString string) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredUnless,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAny[S ~[]T, T any](names string, required bool) yav.ValidateFunc[S] {
	if !required {
		return OmitEmpty[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAny[S ~[]T, T any](names string, required bool) yav.ValidateFunc[S] {
	if !required {
		return OmitEmpty[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAll[S ~[]T, T any](names string, required bool) yav.ValidateFunc[S] {
	if !required {
		return OmitEmpty[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAll[S ~[]T, T any](names string, required bool) yav.ValidateFunc[S] {
	if !required {
		return OmitEmpty[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
