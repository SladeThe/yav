package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
)

type Element generic.Type

var (
	requiredWithAnyElementFuncs map[string]yav.ValidateFunc[Element]
)

func RequiredWithAnyElement() accumulators.RequiredWithAny[Element] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyElement)
}

func provideRequiredWithAnyElement(names string, required bool) yav.ValidateFunc[Element] {
	if !required {
		return OmitEmpty[Element]
	}

	if validateFunc, ok := requiredWithAnyElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyElementFuncs, names, requiredWithAnyElement(names))
}

func requiredWithAnyElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
