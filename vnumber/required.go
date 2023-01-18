package vnumber

import (
	"golang.org/x/exp/constraints"

	"github.com/SladeThe/yav"
)

func OmitEmpty[T constraints.Integer | constraints.Float](_ string, value T) (stop bool, err error) {
	return value == 0, nil
}

func Required[T constraints.Integer | constraints.Float](name string, value T) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}
