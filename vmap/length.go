package vmap

import (
	"reflect"
	"strconv"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

type key[P comparable] struct {
	mapPackage   string
	mapType      string
	keyPackage   string
	keyType      string
	valuePackage string
	valueType    string
	parameter    P
}

func newKey[P comparable, M ~map[K]V, K comparable, V any](parameter P) key[P] {
	mapType := reflect.TypeOf(M(nil))

	keyType, valueType := mapType.Key(), mapType.Elem()

	return key[P]{
		mapPackage:   mapType.PkgPath(),
		mapType:      mapType.String(),
		keyPackage:   keyType.PkgPath(),
		keyType:      keyType.String(),
		valuePackage: valueType.PkgPath(),
		valueType:    valueType.String(),
		parameter:    parameter,
	}
}

type betweenKey struct {
	min, max int
}

var (
	minFuncs     map[key[int]]any
	maxFuncs     map[key[int]]any
	betweenFuncs map[key[betweenKey]]any
)

func Min[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	k := newKey[int, M](parameter)

	if validateFunc, ok := minFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[int], any](&minFuncs, k, min[M](parameter)).(yav.ValidateFunc[M])
}

func Max[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	k := newKey[int, M](parameter)

	if validateFunc, ok := maxFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[int], any](&maxFuncs, k, max[M](parameter)).(yav.ValidateFunc[M])
}

func Between[M ~map[K]V, K comparable, V any](min, max int) yav.ValidateFunc[M] {
	k := newKey[betweenKey, M](betweenKey{min: min, max: max})

	if validateFunc, ok := betweenFuncs[k]; ok {
		return validateFunc.(yav.ValidateFunc[M])
	}

	return internal.RegisterMapEntry[key[betweenKey], any](&betweenFuncs, k, between[M](min, max)).(yav.ValidateFunc[M])
}

func min[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) < parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func max[M ~map[K]V, K comparable, V any](parameter int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) > parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(parameter),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func between[M ~map[K]V, K comparable, V any](min, max int) yav.ValidateFunc[M] {
	return func(name string, value M) (stop bool, err error) {
		if len(value) < min {
			return true, yav.Error{
				CheckName: yav.CheckNameMin,
				Parameter: strconv.Itoa(min),
				ValueName: name,
				Value:     value,
			}
		}

		if len(value) > max {
			return true, yav.Error{
				CheckName: yav.CheckNameMax,
				Parameter: strconv.Itoa(max),
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
