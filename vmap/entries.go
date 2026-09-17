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
		yavErrs.AppendConverted(err, withKeyName(name, key, err))
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
		yavErrs.AppendConverted(err, withKeyName(name, key, err))
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
		yavErrs.AppendConverted(err, withKeyName(name, key, err))
		if stop {
			break
		}
	}
}

func withKeyName[K any](name string, key K, err error) func(yav.Error) yav.Error {
	if err == nil {
		return nil
	}

	switch typedErr := err.(type) {
	case yav.Error:
	case yav.Errors:
		if len(typedErr.Validation) == 0 {
			return nil
		}
	default:
		return nil
	}

	return withKeyNameYAV(name, key)
}

func withKeyNameYAV[K any](name string, key K) func(yav.Error) yav.Error {
	return func(yavErr yav.Error) yav.Error {
		if !strings.HasPrefix(yavErr.ValueName, name) {
			return yavErr
		}

		suffix := yavErr.ValueName[len(name):]

		if suffix != "" && suffix[0] != '.' && suffix[0] != '[' {
			if name != "" {
				return yavErr
			}

			yavErr.ValueName = fmt.Sprintf("[%v].%s", key, suffix)
		} else {
			yavErr.ValueName = fmt.Sprintf("%s[%v]%s", name, key, suffix)
		}

		return yavErr
	}
}
