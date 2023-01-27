package vmap

import (
	"fmt"
	"strings"

	"go.uber.org/multierr"

	"github.com/SladeThe/yav"
)

func Keys[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[K]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		for k := range value {
			multierr.AppendInto(&err, keyChain(name, k, validateFuncs))
		}

		return
	}
}

func Values[M ~map[K]V, K comparable, V any](validateFuncs ...yav.ValidateFunc[V]) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		for k, v := range value {
			multierr.AppendInto(&err, valueChain(name, k, v, validateFuncs))
		}

		return
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

	if validationErr, ok := err.(yav.Error); ok {
		if !strings.HasPrefix(validationErr.ValueName, name) {
			return err
		}

		if len(validationErr.ValueName) > len(name) && validationErr.ValueName[len(name)] != '.' {
			validationErr.ValueName = fmt.Sprintf("%s[%v].%s", name, key, validationErr.ValueName[len(name):])
		} else {
			validationErr.ValueName = fmt.Sprintf("%s[%v]%s", name, key, validationErr.ValueName[len(name):])
		}

		return validationErr
	}

	partialErrs := multierr.Errors(err)
	if len(partialErrs) <= 1 {
		return err
	}

	var combinedErr error

	for _, partialErr := range partialErrs {
		multierr.AppendInto(&combinedErr, withKeyName(name, key, partialErr))
	}

	return combinedErr
}
