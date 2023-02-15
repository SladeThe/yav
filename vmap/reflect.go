package vmap

import (
	"reflect"
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
