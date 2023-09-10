package vmap

import (
	"fmt"
	"strings"

	"github.com/SladeThe/yav"
)

func Keys[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[K]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		var yavErrs yav.Errors

		for k := range value {
			yavErrs.Append(keyChain(name, k, validateFuncs))
		}

		return false, yavErrs.AsError()
	}
}

func Values[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[V]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		var yavErrs yav.Errors

		for k, v := range value {
			yavErrs.Append(valueChain(name, k, v, validateFuncs))
		}

		return false, yavErrs.AsError()
	}
}

func keyChain[K comparable](name string, key K, validateFuncs []yav.ValidateFunc[K]) error {
	for _, validateFunc := range validateFuncs {
		if stop, err := validateFunc(name, key); stop {
			return withKeyName(name, key, err)
		}
	}

	return nil
}

func valueChain[K comparable, V any](name string, key K, value V, validateFuncs []yav.ValidateFunc[V]) error {
	for _, validateFunc := range validateFuncs {
		if stop, err := validateFunc(name, value); stop {
			return withKeyName(name, key, err)
		}
	}

	return nil
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
