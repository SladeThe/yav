package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	requiredWithoutAllElementFuncs map[string]yav.ValidateFunc[Element]
)

func RequiredWithoutAllElement() accumulators.RequiredWithoutAll[Element] {
	return accumulators.NewRequiredWithoutAll(provideRequiredWithoutAllElement)
}

func provideRequiredWithoutAllElement(names string, required bool) yav.ValidateFunc[Element] {
	if !required {
		return OmitEmpty[Element]
	}

	if validateFunc, ok := requiredWithoutAllElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAllElementFuncs, names, requiredWithoutAllElement(names))
}

func requiredWithoutAllElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
