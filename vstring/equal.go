package vstring

import (
	"strings"

	"github.com/SladeThe/yav"
)

func Equal(parameter string) yav.ValidateFunc[string] {
	return equal(parameter).validate
}

func OneOf(parameters ...string) yav.ValidateFunc[string] {
	return oneOf(parameters).validate
}

type equal string

func (e equal) validate(name string, value string) (stop bool, err error) {
	if value != string(e) {
		return false, yav.Error{
			CheckName: yav.CheckNameEqual,
			Parameter: string(e),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type oneOf []string

func (o oneOf) validate(name string, value string) (stop bool, err error) {
	for _, parameter := range o {
		if value == parameter {
			return false, nil
		}
	}

	return false, yav.Error{
		CheckName: yav.CheckNameOneOf,
		Parameter: strings.Join(o, " "),
		ValueName: name,
		Value:     value,
	}
}
