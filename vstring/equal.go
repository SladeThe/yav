package vstring

import (
	"strings"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

var (
	equalFuncs = make(map[string]yav.ValidateFunc[string])
	oneOfFuncs = make(map[string]yav.ValidateFunc[string])
)

func Equal(parameter string) yav.ValidateFunc[string] {
	if validateFunc, ok := equalFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterValidateFunc(&equalFuncs, parameter, equal(parameter))
}

func equal(parameter string) yav.ValidateFunc[string] {
	return func(name string, value string) (stop bool, err error) {
		if value != parameter {
			return false, yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: parameter,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}

func OneOf(parameter string) yav.ValidateFunc[string] {
	if validateFunc, ok := oneOfFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterValidateFunc(&oneOfFuncs, parameter, oneOf(parameter))
}

func oneOf(parameter string) yav.ValidateFunc[string] {
	var expectedValues []string

	for _, expectedValue := range strings.Split(parameter, " ") {
		if expectedValue = strings.TrimSpace(expectedValue); expectedValue != "" {
			expectedValues = append(expectedValues, expectedValue)
		}
	}

	expectedValues = expectedValues[:len(expectedValues):len(expectedValues)]

	return func(name string, value string) (stop bool, err error) {
		for _, expectedValue := range expectedValues {
			if value == expectedValue {
				return false, nil
			}
		}

		return false, yav.Error{
			CheckName: yav.CheckNameOneOf,
			Parameter: parameter,
			ValueName: name,
			Value:     value,
		}
	}
}
