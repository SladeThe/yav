package vmap

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

func OmitEmpty[M ~map[K]V, K comparable, V any](_ string, value M) (stop bool, err error) {
	return len(value) == 0, nil
}

func Required[M ~map[K]V, K comparable, V any](name string, value M) (stop bool, err error) {
	if len(value) == 0 {
		return false, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredWithAny[M ~map[K]V, K comparable, V any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[M] {
	if !accumulator.IsEnabled() {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
		if len(value) == 0 {
			return false, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithoutAny[M ~map[K]V, K comparable, V any](
	fields string,
	accumulator yav.Accumulator,
) yav.ValidateFunc[M] {
	if !accumulator.IsEnabled() {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
		if len(value) == 0 {
			return false, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithAll[M ~map[K]V, K comparable, V any](fields string, accumulator yav.Accumulator) yav.ValidateFunc[M] {
	if !accumulator.IsEnabled() {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
		if len(value) == 0 {
			return false, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func RequiredWithoutAll[M ~map[K]V, K comparable, V any](
	fields string,
	accumulator yav.Accumulator,
) yav.ValidateFunc[M] {
	if !accumulator.IsEnabled() {
		return internal.Valid[M]
	}

	// TODO avoid allocations ?

	return func(name string, value M) (stop bool, err error) {
		if len(value) == 0 {
			return false, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: fields,
				ValueName: name,
			}
		}

		return false, nil
	}
}
