package vstring

import (
	"strconv"

	"github.com/SladeThe/yav"
)

func Min(parameter int) yav.ValidateFunc[string] {
	return minimum(parameter).validate
}

func Max(parameter int) yav.ValidateFunc[string] {
	return maximum(parameter).validate
}

func Between(minimum, maximum int) yav.ValidateFunc[string] {
	return between{min: minimum, max: maximum}.validate
}

type minimum int

func (m minimum) validate(name string, value string) (stop bool, err error) {
	if len(value) < int(m) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: strconv.Itoa(int(m)),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type maximum int

func (m maximum) validate(name string, value string) (stop bool, err error) {
	if len(value) > int(m) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: strconv.Itoa(int(m)),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type between struct {
	min, max int
}

func (b between) validate(name string, value string) (stop bool, err error) {
	if len(value) < b.min {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: strconv.Itoa(b.min),
			ValueName: name,
			Value:     value,
		}
	}

	if len(value) > b.max {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: strconv.Itoa(b.max),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
