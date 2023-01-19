package vnumber

import (
	"fmt"

	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	equalElementFuncs map[Element]yav.ValidateFunc[Element]
)

func EqualElement(parameter Element) yav.ValidateFunc[Element] {
	if validateFunc, ok := equalElementFuncs[parameter]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&equalElementFuncs, parameter, equalElement(parameter))
}

func OneOfElement(parameters ...Element) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		for _, parameter := range parameters {
			if value == parameter {
				return false, nil
			}
		}

		var parameterBuilder strings.Builder

		for i, parameter := range parameters {
			if i == 0 {
				parameterBuilder.WriteString(fmt.Sprintf("%v", parameter))
			} else {
				parameterBuilder.WriteString(fmt.Sprintf(" %v", parameter))
			}
		}

		return true, yav.Error{
			CheckName: yav.CheckNameOneOf,
			Parameter: parameterBuilder.String(),
			ValueName: name,
			Value:     value,
		}
	}
}

func equalElement(parameter Element) yav.ValidateFunc[Element] {
	parameterString := fmt.Sprintf("%v", parameter)

	return func(name string, value Element) (stop bool, err error) {
		if value != parameter {
			return true, yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: parameterString,
				ValueName: name,
				Value:     value,
			}
		}

		return false, nil
	}
}
