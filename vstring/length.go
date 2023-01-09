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

type min int

func (m min) validate(name string, value string) (stop bool, err error) {
	if len(value) < int(m) {
		return false, yav.Error{
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
		return false, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: strconv.Itoa(int(m)),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
