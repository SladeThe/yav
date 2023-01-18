package vmap

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
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
