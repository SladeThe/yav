package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	excludedWithAnyElementFuncs map[string]yav.ValidateFunc[Element]
)

func ExcludedWithAnyElement() accumulators.ExcludedWithAny[Element] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyElement)
}

func provideExcludedWithAnyElement(names string, excluded bool) yav.ValidateFunc[Element] {
	if !excluded {
		return yav.Next[Element]
	}

	if validateFunc, ok := excludedWithAnyElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyElementFuncs, names, excludedWithAnyElement(names))
}

func excludedWithAnyElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
