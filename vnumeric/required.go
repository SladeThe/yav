package vnumeric

import (
	"golang.org/x/exp/constraints"

	"github.com/SladeThe/yav"
)

func OmitEmpty[T constraints.Integer](_ string, value T) (stop bool, err error) {
	return value == 0, nil
}

func Required[T constraints.Integer](name string, value T) (stop bool, err error) {
	if value == 0 {
		return false, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}
