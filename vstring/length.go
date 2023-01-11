package vstring

import (
	"strconv"

	"github.com/SladeThe/yav"
)

func Min(parameter int) yav.ValidateFunc[string] {
	return min(parameter).validate
}

func Max(parameter int) yav.ValidateFunc[string] {
	return max(parameter).validate
}

func Between(min, max int) yav.ValidateFunc[string] {
	return between{min: min, max: max}.validate
}

type min int

func (m min) validate(name string, value string) (stop bool, err error) {
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

type max int

func (m max) validate(name string, value string) (stop bool, err error) {
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
