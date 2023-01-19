package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	requiredWithAllElementFuncs map[string]yav.ValidateFunc[Element]
)

func RequiredWithAllElement() accumulators.RequiredWithAll[Element] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllElement)
}

func provideRequiredWithAllElement(names string, required bool) yav.ValidateFunc[Element] {
	if !required {
		return OmitEmpty[Element]
	}

	if validateFunc, ok := requiredWithAllElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllElementFuncs, names, requiredWithAllElement(names))
}

func requiredWithAllElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
