package vtime

import (
	"time"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredWithAnyFuncs    map[string]yav.ValidateFunc[time.Time]
	requiredWithoutAnyFuncs map[string]yav.ValidateFunc[time.Time]
	requiredWithAllFuncs    map[string]yav.ValidateFunc[time.Time]
	requiredWithoutAllFuncs map[string]yav.ValidateFunc[time.Time]
)

func OmitEmpty(_ string, value time.Time) (stop bool, err error) {
	return value.IsZero(), nil
}

func Required(name string, value time.Time) (stop bool, err error) {
	if value.IsZero() {
		return true, yav.ErrRequired(name)
	}

	return false, nil
}

func RequiredIf(conditionString string, condition bool) yav.ValidateFunc[time.Time] {
	if !condition {
		return OmitEmpty
	}

	return requiredIf(conditionString).validate
}

func RequiredUnless(conditionString string, condition bool) yav.ValidateFunc[time.Time] {
	if condition {
		return OmitEmpty
	}

	return requiredUnless(conditionString).validate
}

func RequiredWithAny() accumulators.RequiredWithAny[time.Time] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAny)
}

func RequiredWithoutAny() accumulators.RequiredWithoutAny[time.Time] {
	return accumulators.NewRequiredWithoutAny(provideRequiredWithoutAny)
}

func RequiredWithAll() accumulators.RequiredWithAll[time.Time] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAll)
}

func RequiredWithoutAll() accumulators.RequiredWithoutAll[time.Time] {
	return accumulators.NewRequiredWithoutAll(provideRequiredWithoutAll)
}

type requiredIf string

func (r requiredIf) validate(name string, value time.Time) (stop bool, err error) {
	if value.IsZero() {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

type requiredUnless string

func (r requiredUnless) validate(name string, value time.Time) (stop bool, err error) {
	if value.IsZero() {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredUnless,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func provideRequiredWithAny(names string, required bool) yav.ValidateFunc[time.Time] {
	if !required {
		return OmitEmpty
	}

	if validateFunc, ok := requiredWithAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyFuncs, names, requiredWithAny(names))
}

func requiredWithAny(names string) yav.ValidateFunc[time.Time] {
	return func(name string, value time.Time) (stop bool, err error) {
		if value.IsZero() {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAny(names string, required bool) yav.ValidateFunc[time.Time] {
	if !required {
		return OmitEmpty
	}

	if validateFunc, ok := requiredWithoutAnyFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAnyFuncs, names, requiredWithoutAny(names))
}

func requiredWithoutAny(names string) yav.ValidateFunc[time.Time] {
	return func(name string, value time.Time) (stop bool, err error) {
		if value.IsZero() {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAll(names string, required bool) yav.ValidateFunc[time.Time] {
	if !required {
		return OmitEmpty
	}

	if validateFunc, ok := requiredWithAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllFuncs, names, requiredWithAll(names))
}

func requiredWithAll(names string) yav.ValidateFunc[time.Time] {
	return func(name string, value time.Time) (stop bool, err error) {
		if value.IsZero() {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAll(names string, required bool) yav.ValidateFunc[time.Time] {
	if !required {
		return OmitEmpty
	}

	if validateFunc, ok := requiredWithoutAllFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAllFuncs, names, requiredWithoutAll(names))
}

func requiredWithoutAll(names string) yav.ValidateFunc[time.Time] {
	return func(name string, value time.Time) (stop bool, err error) {
		if value.IsZero() {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
