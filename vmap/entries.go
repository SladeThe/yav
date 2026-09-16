package vmap

import (
	"fmt"
	"strings"

	"github.com/SladeThe/yav"
)

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Keys[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[K]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		var yavErrs yav.Errors

		for k := range value {
			keyChain(&yavErrs, name, k, validateFuncs)
		}

		return false, yavErrs.AsError()
	}
}

func Values[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[V]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		var yavErrs yav.Errors

		for k, v := range value {
			valueChain(&yavErrs, name, k, v, validateFuncs)
		}

		return false, yavErrs.AsError()
	}
}

func Entries[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[Entry[K, V]]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		var yavErrs yav.Errors

		for k, v := range value {
			entryChain(&yavErrs, name, k, v, validateFuncs)
		}

		return false, yavErrs.AsError()
	}
}

func keyChain[K comparable](yavErrs *yav.Errors, name string, key K, validateFuncs []yav.ValidateFunc[K]) {
	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, key)
		yavErrs.Append(withKeyName(name, key, err))
		if stop {
			break
		}
	}
}

func valueChain[K comparable, V any](
	yavErrs *yav.Errors, name string, key K, value V, validateFuncs []yav.ValidateFunc[V],
) {
	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, value)
		yavErrs.Append(withKeyName(name, key, err))
		if stop {
			break
		}
	}
}

func entryChain[K comparable, V any](
	yavErrs *yav.Errors, name string, key K, value V, validateFuncs []yav.ValidateFunc[Entry[K, V]],
) {
	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, Entry[K, V]{key, value})
		yavErrs.Append(withKeyName(name, key, err))
		if stop {
			break
		}
	}
}

func withKeyName[K any](name string, key K, err error) error {
	if err == nil {
		return nil
	}

	switch typedErr := err.(type) {
	case yav.Error:
		return withKeyNameYAV(name, key, typedErr)
	case yav.Errors:
		for i, yavErr := range typedErr.Validation {
			typedErr.Validation[i] = withKeyNameYAV(name, key, yavErr)
		}

		return typedErr
	default:
		return err
	}
}

func withKeyNameYAV[K any](name string, key K, yavErr yav.Error) yav.Error {
	if !strings.HasPrefix(yavErr.ValueName, name) {
		return yavErr
	}

	if len(yavErr.ValueName) > len(name) && yavErr.ValueName[len(name)] != '.' {
		yavErr.ValueName = fmt.Sprintf("%s[%v].%s", name, key, yavErr.ValueName[len(name):])
	} else {
		yavErr.ValueName = fmt.Sprintf("%s[%v]%s", name, key, yavErr.ValueName[len(name):])
	}

	return yavErr
}
