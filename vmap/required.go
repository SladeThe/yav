package vmap

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
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

func RequiredWithAny[M ~map[K]V, K comparable, V any](fields string) accumulators.RequiredWithAny[M] {
	return accumulators.NewRequiredWithAny(fields, provideRequiredWithAny[M])
}

func RequiredWithoutAny[M ~map[K]V, K comparable, V any](fields string) accumulators.RequiredWithoutAny[M] {
	return accumulators.NewRequiredWithoutAny(fields, provideRequiredWithoutAny[M])
}

func RequiredWithAll[M ~map[K]V, K comparable, V any](fields string) accumulators.RequiredWithAll[M] {
	return accumulators.NewRequiredWithAll(fields, provideRequiredWithAll[M])
}

func RequiredWithoutAll[M ~map[K]V, K comparable, V any](fields string) accumulators.RequiredWithoutAll[M] {
	return accumulators.NewRequiredWithoutAll(fields, provideRequiredWithoutAll[M])
}

func provideRequiredWithAny[M ~map[K]V, K comparable, V any](fields string, enabled bool) yav.ValidateFunc[M] {
	if !enabled {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithoutAny[M ~map[K]V, K comparable, V any](fields string, enabled bool) yav.ValidateFunc[M] {
	if !enabled {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithAll[M ~map[K]V, K comparable, V any](fields string, enabled bool) yav.ValidateFunc[M] {
	if !enabled {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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

func provideRequiredWithoutAll[M ~map[K]V, K comparable, V any](fields string, enabled bool) yav.ValidateFunc[M] {
	if !enabled {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
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
