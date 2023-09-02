package vbytes

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedWithAnyFuncs    map[string]yav.ValidateFunc[[]byte]
	excludedWithoutAnyFuncs map[string]yav.ValidateFunc[[]byte]
	excludedWithAllFuncs    map[string]yav.ValidateFunc[[]byte]
	excludedWithoutAllFuncs map[string]yav.ValidateFunc[[]byte]
)

func ExcludedIf(conditionString string, condition bool) yav.ValidateFunc[[]byte] {
	if !condition {
		return yav.Next[[]byte]
	}

	return excludedIf(conditionString).validate
}

func ExcludedUnless(conditionString string, condition bool) yav.ValidateFunc[[]byte] {
	if condition {
		return yav.Next[[]byte]
	}

	return excludedUnless(conditionString).validate
}

func ExcludedWithAny() accumulators.ExcludedWithAny[[]byte] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny)
}

func ExcludedWithoutAny() accumulators.ExcludedWithoutAny[[]byte] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny)
}

func ExcludedWithAll() accumulators.ExcludedWithAll[[]byte] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll)
}

func ExcludedWithoutAll() accumulators.ExcludedWithoutAll[[]byte] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll)
}

type excludedIf string

func (r excludedIf) validate(name string, value []byte) (stop bool, err error) {
	if len(value) > 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

type excludedUnless string

func (r excludedUnless) validate(name string, value []byte) (stop bool, err error) {
	if len(value) > 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func provideExcludedWithAny(names string, excluded bool) yav.ValidateFunc[[]byte] {
	if !excluded {
		return yav.Next[[]byte]
	}

	if validateFunc, ok := excludedWithAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyFuncs, names, excludedWithAny(names))
}

func excludedWithAny(names string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAny(names string, excluded bool) yav.ValidateFunc[[]byte] {
	if !excluded {
		return yav.Next[[]byte]
	}

	if validateFunc, ok := excludedWithoutAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAnyFuncs, names, excludedWithoutAny(names))
}

func excludedWithoutAny(names string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAll(names string, excluded bool) yav.ValidateFunc[[]byte] {
	if !excluded {
		return yav.Next[[]byte]
	}

	if validateFunc, ok := excludedWithAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAllFuncs, names, excludedWithAll(names))
}

func excludedWithAll(names string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAll(names string, excluded bool) yav.ValidateFunc[[]byte] {
	if !excluded {
		return yav.Next[[]byte]
	}

	if validateFunc, ok := excludedWithoutAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAllFuncs, names, excludedWithoutAll(names))
}

func excludedWithoutAll(names string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) > 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
