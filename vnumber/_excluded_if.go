package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func ExcludedIfElement(conditionString string, condition bool) yav.ValidateFunc[Element] {
	if !condition {
		return yav.Next[Element]
	}

	return excludedIfElement(conditionString).validate
}

type excludedIfElement string

func (r excludedIfElement) validate(name string, value Element) (stop bool, err error) {
	if value != 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}
