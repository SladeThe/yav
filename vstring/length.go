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

func InRange(min, max int) yav.ValidateFunc[string] {
	return inRange{min: min, max: max}.validate
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

type inRange struct {
	min, max int
}

func (l inRange) validate(name string, value string) (stop bool, err error) {
	if len(value) < l.min {
		return false, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: strconv.Itoa(l.min),
			ValueName: name,
			Value:     value,
		}
	}

	if len(value) > l.max {
		return false, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: strconv.Itoa(l.max),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
