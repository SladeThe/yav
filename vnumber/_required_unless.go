package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func RequiredUnlessElement(conditionString string, condition bool) yav.ValidateFunc[Element] {
	if condition {
		return OmitEmpty[Element]
	}

	return requiredUnlessElement(conditionString).validate
}

type requiredUnlessElement string

func (r requiredUnlessElement) validate(name string, value Element) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}
