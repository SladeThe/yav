package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func ExcludedUnlessElement(conditionString string, condition bool) yav.ValidateFunc[Element] {
	if condition {
		return yav.Next[Element]
	}

	return excludedUnlessElement(conditionString).validate
}

type excludedUnlessElement string

func (r excludedUnlessElement) validate(name string, value Element) (stop bool, err error) {
	if value != 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}
