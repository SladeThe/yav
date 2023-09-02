package vslice

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedIfFuncs     map[key[string]]any
	excludedUnlessFuncs map[key[string]]any
)

func ExcludedIf[S ~[]T, T any](conditionString string, condition bool) yav.ValidateFunc[S] {
	if !condition {
		return yav.Next[S]
	}

	k := newKey[string, S](conditionString)

	if validateFunc, ok := excludedIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedIfFuncs, k, excludedIf[S](conditionString),
	).(yav.ValidateFunc[S])
}

func ExcludedUnless[S ~[]T, T any](conditionString string, condition bool) yav.ValidateFunc[S] {
	if condition {
		return yav.Next[S]
	}

	k := newKey[string, S](conditionString)

	if validateFunc, ok := excludedUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[S])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedUnlessFuncs, k, excludedUnless[S](conditionString),
	).(yav.ValidateFunc[S])
}

func ExcludedWithAny[S ~[]T, T any]() accumulators.ExcludedWithAny[S] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny[S])
}

func ExcludedWithoutAny[S ~[]T, T any]() accumulators.ExcludedWithoutAny[S] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny[S])
}

func ExcludedWithAll[S ~[]T, T any]() accumulators.ExcludedWithAll[S] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll[S])
}

func ExcludedWithoutAll[S ~[]T, T any]() accumulators.ExcludedWithoutAll[S] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll[S])
}

func excludedIf[S ~[]T, T any](conditionString string) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedIf,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func excludedUnless[S ~[]T, T any](conditionString string) yav.ValidateFunc[S] {
	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedUnless,
				Parameter: conditionString,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAny[S ~[]T, T any](names string, excluded bool) yav.ValidateFunc[S] {
	if !excluded {
		return yav.Next[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAny[S ~[]T, T any](names string, excluded bool) yav.ValidateFunc[S] {
	if !excluded {
		return yav.Next[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAll[S ~[]T, T any](names string, excluded bool) yav.ValidateFunc[S] {
	if !excluded {
		return yav.Next[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAll[S ~[]T, T any](names string, excluded bool) yav.ValidateFunc[S] {
	if !excluded {
		return yav.Next[S]
	}

	// TODO avoid allocations ?

	return func(name string, value S) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
