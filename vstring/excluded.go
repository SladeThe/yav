package vstring

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	excludedWithAnyFuncs    map[string]yav.ValidateFunc[string]
	excludedWithoutAnyFuncs map[string]yav.ValidateFunc[string]
	excludedWithAllFuncs    map[string]yav.ValidateFunc[string]
	excludedWithoutAllFuncs map[string]yav.ValidateFunc[string]
)

func ExcludedIf(conditionString string, condition bool) yav.ValidateFunc[string] {
	if !condition {
		return yav.Next[string]
	}

	return excludedIf(conditionString).validate
}

func ExcludedUnless(conditionString string, condition bool) yav.ValidateFunc[string] {
	if condition {
		return yav.Next[string]
	}

	return excludedUnless(conditionString).validate
}

func ExcludedWithAny() accumulators.ExcludedWithAny[string] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAny)
}

func ExcludedWithoutAny() accumulators.ExcludedWithoutAny[string] {
	return accumulators.NewExcludedWithoutAny(provideExcludedWithoutAny)
}

func ExcludedWithAll() accumulators.ExcludedWithAll[string] {
	return accumulators.NewExcludedWithAll(provideExcludedWithAll)
}

func ExcludedWithoutAll() accumulators.ExcludedWithoutAll[string] {
	return accumulators.NewExcludedWithoutAll(provideExcludedWithoutAll)
}

type excludedIf string

func (r excludedIf) validate(name string, value string) (stop bool, err error) {
	if value != "" {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

type excludedUnless string

func (r excludedUnless) validate(name string, value string) (stop bool, err error) {
	if value != "" {
		return true, yav.Error{
			CheckName: yav.CheckNameExcludedUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func provideExcludedWithAny(names string, excluded bool) yav.ValidateFunc[string] {
	if !excluded {
		return yav.Next[string]
	}

	if validateFunc, ok := excludedWithAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyFuncs, names, excludedWithAny(names))
}

func excludedWithAny(names string) yav.ValidateFunc[string] {
	return func(name string, value string) (stop bool, err error) {
		if value != "" {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAny(names string, excluded bool) yav.ValidateFunc[string] {
	if !excluded {
		return yav.Next[string]
	}

	if validateFunc, ok := excludedWithoutAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAnyFuncs, names, excludedWithoutAny(names))
}

func excludedWithoutAny(names string) yav.ValidateFunc[string] {
	return func(name string, value string) (stop bool, err error) {
		if value != "" {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithAll(names string, excluded bool) yav.ValidateFunc[string] {
	if !excluded {
		return yav.Next[string]
	}

	if validateFunc, ok := excludedWithAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAllFuncs, names, excludedWithAll(names))
}

func excludedWithAll(names string) yav.ValidateFunc[string] {
	return func(name string, value string) (stop bool, err error) {
		if value != "" {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideExcludedWithoutAll(names string, excluded bool) yav.ValidateFunc[string] {
	if !excluded {
		return yav.Next[string]
	}

	if validateFunc, ok := excludedWithoutAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithoutAllFuncs, names, excludedWithoutAll(names))
}

func excludedWithoutAll(names string) yav.ValidateFunc[string] {
	return func(name string, value string) (stop bool, err error) {
		if value != "" {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
