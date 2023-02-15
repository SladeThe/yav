package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func RequiredIfElement(conditionString string, condition bool) yav.ValidateFunc[Element] {
	if !condition {
		return OmitEmpty[Element]
	}

	return requiredIfElement(conditionString).validate
}

type requiredIfElement string

func (r requiredIfElement) validate(name string, value Element) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}
