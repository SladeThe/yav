package vmap

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredIfFuncs     map[key[string]]any
	requiredUnlessFuncs map[key[string]]any
)

func OmitEmpty[M ~map[K]V, K comparable, V any](_ string, value M) (stop bool, err error) {
	return len(value) == 0, nil
}

func Required[M ~map[K]V, K comparable, V any](name string, value M) (stop bool, err error) {
	if len(value) == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIf[M ~map[K]V, K comparable, V any](conditionString string, condition bool) yav.ValidateFunc[M] {
	if !condition {
		return OmitEmpty[M]
	}

	k := newKey[string, M](conditionString)

	if validateFunc, ok := requiredIfFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredIfFuncs, k, requiredIf[M](conditionString),
	).(yav.ValidateFunc[M])
}

func RequiredUnless[M ~map[K]V, K comparable, V any](conditionString string, condition bool) yav.ValidateFunc[M] {
	if condition {
		return OmitEmpty[M]
	}

	k := newKey[string, M](conditionString)

	if validateFunc, ok := requiredUnlessFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[string], any](
		&requiredUnlessFuncs, k, requiredUnless[M](conditionString),
	).(yav.ValidateFunc[M])
}

func RequiredWithAny[M ~map[K]V, K comparable, V any]() accumulators.RequiredWithAny[M] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAny[M])
}

func RequiredWithoutAny[M ~map[K]V, K comparable, V any]() accumulators.RequiredWithoutAny[M] {
	return accumulators.NewRequiredWithoutAny(provideRequiredWithoutAny[M])
}

func RequiredWithAll[M ~map[K]V, K comparable, V any]() accumulators.RequiredWithAll[M] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAll[M])
}

func RequiredWithoutAll[M ~map[K]V, K comparable, V any]() accumulators.RequiredWithoutAll[M] {
	return accumulators.NewRequiredWithoutAll(provideRequiredWithoutAll[M])
}

func requiredIf[M ~map[K]V, K comparable, V any](conditionString string) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
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

func requiredUnless[M ~map[K]V, K comparable, V any](conditionString string) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithAny[M ~map[K]V, K comparable, V any](names string, required bool) yav.ValidateFunc[M] {
	if !required {
		return OmitEmpty[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithoutAny[M ~map[K]V, K comparable, V any](names string, required bool) yav.ValidateFunc[M] {
	if !required {
		return OmitEmpty[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithAll[M ~map[K]V, K comparable, V any](names string, required bool) yav.ValidateFunc[M] {
	if !required {
		return OmitEmpty[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithoutAll[M ~map[K]V, K comparable, V any](names string, required bool) yav.ValidateFunc[M] {
	if !required {
		return OmitEmpty[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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
