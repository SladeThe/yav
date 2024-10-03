package vpointer

import (
	"github.com/SladeThe/yav"
)

func OmitEmpty(_ string, value any) (stop bool, err error) {
	return value == nil, nil
}

func Required(name string, value any) (stop bool, err error) {
	if value == nil {
		return true, yav.ErrRequired(name)
	}

	return false, nil
}
