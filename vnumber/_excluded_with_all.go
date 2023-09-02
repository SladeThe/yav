package vnumber

import (
	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

type Element generic.Type

var (
	excludedWithAllElementFuncs map[string]yav.ValidateFunc[Element]
)

func ExcludedWithAllElement() accumulators.ExcludedWithAll[Element] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAllElement)
}

func provideExcludedWithAllElement(names string, excluded bool) yav.ValidateFunc[Element] {
	if !excluded {
		return yav.Next[Element]
	}

	if validateFunc, ok := excludedWithAllElementFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAllElementFuncs, names, excludedWithAllElement(names))
}

func excludedWithAllElement(names string) yav.ValidateFunc[Element] {
	return func(name string, value Element) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
