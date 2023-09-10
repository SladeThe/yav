package vslice

import (
	"reflect"
)

type key[P comparable] struct {
	sliceType reflect.Type
	parameter P
}

func newKey[P comparable, S ~[]T, T any](parameter P) key[P] {
	return key[P]{
		sliceType: reflect.TypeOf(S(nil)),
		parameter: parameter,
	}
}
