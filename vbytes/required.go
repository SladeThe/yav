package vbytes

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredWithAnyFuncs    map[string]yav.ValidateFunc[[]byte]
	requiredWithoutAnyFuncs map[string]yav.ValidateFunc[[]byte]
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

func RequiredWithAny(fields string, accumulator yav.Accumulator) yav.ValidateFunc[[]byte] {
	if !accumulator.IsEnabled() {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithAnyFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyFuncs, fields, requiredWithAny(fields))
}

func RequiredWithoutAny(fields string, accumulator yav.Accumulator) yav.ValidateFunc[[]byte] {
	if !accumulator.IsEnabled() {
		return internal.Valid[[]byte]
	}

	if validateFunc, ok := requiredWithoutAnyFuncs[fields]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithoutAnyFuncs, fields, requiredWithoutAny(fields))
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
