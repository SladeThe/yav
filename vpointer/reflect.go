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
		pointedType: reflect.TypeFor[T](),
		parameter:   parameter,
	}
}
