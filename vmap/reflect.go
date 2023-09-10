package vmap

import (
	"reflect"
)

type key[P comparable] struct {
	mapType   reflect.Type
	parameter P
}

func newKey[P comparable, M ~map[K]V, K comparable, V any](parameter P) key[P] {
	return key[P]{
		mapType:   reflect.TypeOf(M(nil)),
		parameter: parameter,
	}
}
