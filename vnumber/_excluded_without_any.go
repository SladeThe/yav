package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	excludedWithoutAnyElementFuncs map[string]yav.ValidateFunc[Element]
)

func ExcludedWithoutAnyElement() accumulators.ExcludedWithoutAny[Element] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAnyElement)
}

func provideExcludedWithoutAnyElement(names string, excluded bool) yav.ValidateFunc[Element] {
	if !excluded {
		return yav.Next[Element]
	}

	if validateFunc, ok := excludedWithoutAnyElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAnyElementFuncs, names, excludedWithoutAnyElement(names))
}

func excludedWithoutAnyElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
