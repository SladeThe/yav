package vpointer

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedIfFuncs         map[key[string]]any
	excludedUnlessFuncs     map[key[string]]any
	excludedWithAnyFuncs    map[key[string]]any
	excludedWithoutAnyFuncs map[key[string]]any
	excludedWithAllFuncs    map[key[string]]any
	excludedWithoutAllFuncs map[key[string]]any
)

func ExcludedIf[T any](conditionString string, condition bool) yav.ValidateFunc[*T] {
	if !condition {
		return yav.Next[*T]
	}

	k := newKey[string, T](conditionString)

	if validateFunc, ok := excludedIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedIfFuncs, k, excludedIf[T](conditionString),
	).(yav.ValidateFunc[*T])
}

func ExcludedUnless[T any](conditionString string, condition bool) yav.ValidateFunc[*T] {
	if condition {
		return yav.Next[*T]
	}

	k := newKey[string, T](conditionString)

	if validateFunc, ok := excludedUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedUnlessFuncs, k, excludedUnless[T](conditionString),
	).(yav.ValidateFunc[*T])
}

func ExcludedWithAny[T any]() accumulators.ExcludedWithAny[*T] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny[T])
}

func ExcludedWithoutAny[T any]() accumulators.ExcludedWithoutAny[*T] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny[T])
}

func ExcludedWithAll[T any]() accumulators.ExcludedWithAll[*T] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll[T])
}

func ExcludedWithoutAll[T any]() accumulators.ExcludedWithoutAll[*T] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll[T])
}

func excludedIf[T any](conditionString string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedIf,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func excludedUnless[T any](conditionString string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedUnless,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAny[T any](names string, excluded bool) yav.ValidateFunc[*T] {
	if !excluded {
		return yav.Next[*T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := excludedWithAnyFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedWithAnyFuncs, k, excludedWithAny[T](names),
	).(yav.ValidateFunc[*T])
}

func excludedWithAny[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAny[T any](names string, excluded bool) yav.ValidateFunc[*T] {
	if !excluded {
		return yav.Next[*T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := excludedWithoutAnyFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedWithoutAnyFuncs, k, excludedWithoutAny[T](names),
	).(yav.ValidateFunc[*T])
}

func excludedWithoutAny[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAll[T any](names string, excluded bool) yav.ValidateFunc[*T] {
	if !excluded {
		return yav.Next[*T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := excludedWithAllFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedWithAllFuncs, k, excludedWithAll[T](names),
	).(yav.ValidateFunc[*T])
}

func excludedWithAll[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAll[T any](names string, excluded bool) yav.ValidateFunc[*T] {
	if !excluded {
		return yav.Next[*T]
	}

	k := newKey[string, T](names)

	if validateFunc, ok := excludedWithoutAllFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[*T])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedWithoutAllFuncs, k, excludedWithoutAll[T](names),
	).(yav.ValidateFunc[*T])
}

func excludedWithoutAll[T any](names string) yav.ValidateFunc[*T] {
	return func(name string, value *T) (stop bool, err error) {
		if value != nil {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
