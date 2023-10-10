package vpointer

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredIfFuncs         map[key[string]]any
	requiredUnlessFuncs     map[key[string]]any
	requiredWithAnyFuncs    map[key[string]]any
	requiredWithoutAnyFuncs map[key[string]]any
	requiredWithAllFuncs    map[key[string]]any
	requiredWithoutAllFuncs map[key[string]]any
)

func OmitEmpty[T any](_ string, value *T) (stop bool, err error) {
	return value == nil, nil
}

func Required[T any](name string, value *T) (stop bool, err error) {
	if value == nil {
		return true, yav.ErrRequired(name)
	}

	return false, nil
}

func RequiredIf[T any](conditionString string, condition bool) yav.ValidateFunc[*T] {
	if !condition {
		return OmitEmpty[T]
	}

	k := newKey[string, T](conditionString)

	if validateFunc, ok := requiredIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredIfFuncs, k, requiredIf[T](conditionString),
	).(yav.ValidateFunc[*T])
}

func RequiredUnless[T any](conditionString string, condition bool) yav.ValidateFunc[*T] {
	if condition {
		return OmitEmpty[T]
	}

	k := newKey[string, T](conditionString)

	if validateFunc, ok := requiredUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredUnlessFuncs, k, requiredUnless[T](conditionString),
	).(yav.ValidateFunc[*T])
}

func RequiredWithAny[T any]() accumulators.RequiredWithAny[*T] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAny[T])
}

func RequiredWithoutAny[T any]() accumulators.RequiredWithoutAny[*T] {
	return accumulators.NewRequiredWithoutAny(provideRequiredWithoutAny[T])
}

func RequiredWithAll[T any]() accumulators.RequiredWithAll[*T] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAll[T])
}

func RequiredWithoutAll[T any]() accumulators.RequiredWithoutAll[*T] {
	return accumulators.NewRequiredWithoutAll(provideRequiredWithoutAll[T])
}

func requiredIf[T any](conditionString string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredIf,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func requiredUnless[T any](conditionString string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredUnless,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAny[T any](names string, required bool) yav.ValidateFunc[*T] {
	if !required {
		return OmitEmpty[T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := requiredWithAnyFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredWithAnyFuncs, k, requiredWithAny[T](names),
	).(yav.ValidateFunc[*T])
}

func requiredWithAny[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAny[T any](names string, required bool) yav.ValidateFunc[*T] {
	if !required {
		return OmitEmpty[T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := requiredWithoutAnyFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredWithoutAnyFuncs, k, requiredWithoutAny[T](names),
	).(yav.ValidateFunc[*T])
}

func requiredWithoutAny[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAll[T any](names string, required bool) yav.ValidateFunc[*T] {
	if !required {
		return OmitEmpty[T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := requiredWithAllFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredWithAllFuncs, k, requiredWithAll[T](names),
	).(yav.ValidateFunc[*T])
}

func requiredWithAll[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAll[T any](names string, required bool) yav.ValidateFunc[*T] {
	if !required {
		return OmitEmpty[T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := requiredWithoutAllFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredWithoutAllFuncs, k, requiredWithoutAll[T](names),
	).(yav.ValidateFunc[*T])
}

func requiredWithoutAll[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value == nil {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
