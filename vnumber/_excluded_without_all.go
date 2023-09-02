package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	excludedWithoutAllElementFuncs map[string]yav.ValidateFunc[Element]
)

func ExcludedWithoutAllElement() accumulators.ExcludedWithoutAll[Element] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAllElement)
}

func provideExcludedWithoutAllElement(names string, excluded bool) yav.ValidateFunc[Element] {
	if !excluded {
		return yav.Next[Element]
	}

	if validateFunc, ok := excludedWithoutAllElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAllElementFuncs, names, excludedWithoutAllElement(names))
}

func excludedWithoutAllElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
