package vpointer

import (
	"github.com/SladeThe/yav"
)

func OmitEmpty[T any](_ string, value *T) (stop bool, err error) {
	return value == nil, nil
}

func Required[T any](name string, value *T) (stop bool, err error) {
	if value == nil {
		return true, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}
