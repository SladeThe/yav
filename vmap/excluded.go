package vmap

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedIfFuncs     map[key[string]]any
	excludedUnlessFuncs map[key[string]]any
)

func ExcludedIf[M ~map[K]V, K comparable, V any](conditionString string, condition bool) yav.ValidateFunc[M] {
	if !condition {
		return yav.Next[M]
	}

	k := newKey[string, M](conditionString)

	if validateFunc, ok := excludedIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedIfFuncs, k, excludedIf[M](conditionString),
	).(yav.ValidateFunc[M])
}

func ExcludedUnless[M ~map[K]V, K comparable, V any](conditionString string, condition bool) yav.ValidateFunc[M] {
	if condition {
		return yav.Next[M]
	}

	k := newKey[string, M](conditionString)

	if validateFunc, ok := excludedUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[string], any](
		&excludedUnlessFuncs, k, excludedUnless[M](conditionString),
	).(yav.ValidateFunc[M])
}

func ExcludedWithAny[M ~map[K]V, K comparable, V any]() accumulators.ExcludedWithAny[M] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny[M])
}

func ExcludedWithoutAny[M ~map[K]V, K comparable, V any]() accumulators.ExcludedWithoutAny[M] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny[M])
}

func ExcludedWithAll[M ~map[K]V, K comparable, V any]() accumulators.ExcludedWithAll[M] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll[M])
}

func ExcludedWithoutAll[M ~map[K]V, K comparable, V any]() accumulators.ExcludedWithoutAll[M] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll[M])
}

func excludedIf[M ~map[K]V, K comparable, V any](conditionString string) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
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

func excludedUnless[M ~map[K]V, K comparable, V any](conditionString string) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
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

func provideExcludedWithAny[M ~map[K]V, K comparable, V any](names string, excluded bool) yav.ValidateFunc[M] {
	if !excluded {
		return yav.Next[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideExcludedWithoutAny[M ~map[K]V, K comparable, V any](names string, excluded bool) yav.ValidateFunc[M] {
	if !excluded {
		return yav.Next[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideExcludedWithAll[M ~map[K]V, K comparable, V any](names string, excluded bool) yav.ValidateFunc[M] {
	if !excluded {
		return yav.Next[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideExcludedWithoutAll[M ~map[K]V, K comparable, V any](names string, excluded bool) yav.ValidateFunc[M] {
	if !excluded {
		return yav.Next[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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
