package vpointer

import (
	"reflect"
)

type key[P comparable] struct {
	pointedType reflect.Type
	parameter   P
}

func newKey[P comparable, T any](parameter P) key[P] {
	return key[P]{
		pointedType: reflect.TypeOf((*T)(nil)).Elem(), // TODO TypeFor[T] in Go 1.22
		parameter:   parameter,
	}
}
