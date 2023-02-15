package vslice

import (
	"reflect"
)

type key[P comparable] struct {
	slicePackage string
	sliceType    string
	itemPackage  string
	itemType     string
	parameter    P
}

func newKey[P comparable, S ~[]T, T any](parameter P) key[P] {
	sliceType := reflect.TypeOf(S(nil))

	itemType := sliceType.Elem()

	return key[P]{
		slicePackage: sliceType.PkgPath(),
		sliceType:    sliceType.String(),
		itemPackage:  itemType.PkgPath(),
		itemType:     itemType.String(),
		parameter:    parameter,
	}
}
