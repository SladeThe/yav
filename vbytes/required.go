package vbytes

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredWithAnyFuncs    map[string]yav.ValidateFunc[[]byte]
	requiredWithoutAnyFuncs map[string]yav.ValidateFunc[[]byte]
	requiredWithAllFuncs    map[string]yav.ValidateFunc[[]byte]
	requiredWithoutAllFuncs map[string]yav.ValidateFunc[[]byte]
)

func OmitEmpty(_ string, value []byte) (stop bool, err error) {
	return len(value) == 0, nil
}

func Required(name string, value []byte) (stop bool, err error) {
	if len(value) == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredWithAny(fields string) accumulators.RequiredWithAny[[]byte] {
	return accumulators.NewRequiredWithAny(fields, provideRequiredWithAny)
}

func RequiredWithoutAny(fields string) accumulators.RequiredWithoutAny[[]byte] {
	return accumulators.NewRequiredWithoutAny(fields, provideRequiredWithoutAny)
}

func RequiredWithAll(fields string) accumulators.RequiredWithAll[[]byte] {
	return accumulators.NewRequiredWithAll(fields, provideRequiredWithAll)
}

func RequiredWithoutAll(fields string) accumulators.RequiredWithoutAll[[]byte] {
	return accumulators.NewRequiredWithoutAll(fields, provideRequiredWithoutAll)
}

func provideRequiredWithAny(fields string, enabled bool) yav.ValidateFunc[[]byte] {
	if !enabled {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithAnyFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyFuncs, fields, requiredWithAny(fields))
}

func requiredWithAny(parameter string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: parameter,
				ValueName: name,
			}
		}

		return false, nil
	}
}
func provideRequiredWithoutAny(fields string, enabled bool) yav.ValidateFunc[[]byte] {
	if !enabled {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithoutAnyFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAnyFuncs, fields, requiredWithoutAny(fields))
}

func requiredWithoutAny(parameter string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: parameter,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithAll(fields string, enabled bool) yav.ValidateFunc[[]byte] {
	if !enabled {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithAllFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllFuncs, fields, requiredWithAll(fields))
}

func requiredWithAll(parameter string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: parameter,
				ValueName: name,
			}
		}

		return false, nil
	}
}

func provideRequiredWithoutAll(fields string, enabled bool) yav.ValidateFunc[[]byte] {
	if !enabled {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithoutAllFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAllFuncs, fields, requiredWithoutAll(fields))
}

func requiredWithoutAll(parameter string) yav.ValidateFunc[[]byte] {
	return func(name string, value []byte) (stop bool, err error) {
		if len(value) == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: parameter,
				ValueName: name,
			}
		}

		return false, nil
	}
}
