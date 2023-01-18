package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
)

type Element generic.Type

var (
	requiredWithoutAnyElementFuncs map[string]yav.ValidateFunc[Element]
)

func RequiredWithoutAnyElement() accumulators.RequiredWithoutAny[Element] {
	return accumulators.NewRequiredWithoutAny(provideRequiredWithoutAnyElement)
}

func provideRequiredWithoutAnyElement(names string, required bool) yav.ValidateFunc[Element] {
	if !required {
		return OmitEmpty[Element]
	}

	if validateFunc, ok := requiredWithoutAnyElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAnyElementFuncs, names, requiredWithoutAnyElement(names))
}

func requiredWithoutAnyElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
