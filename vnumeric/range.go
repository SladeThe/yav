package vnumeric

import (
	"golang.org/x/exp/constraints"

	"github.com/SladeThe/yav"
)

func Nonnegative[T constraints.Signed](name string, value T) (stop bool, err error) {
	if value < 0 {
		return false, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: "0",
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Positive[T constraints.Integer](name string, value T) (stop bool, err error) {
	if value < 1 {
		return false, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: "1",
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
