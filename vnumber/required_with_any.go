// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package vnumber

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/accumulators"
	"github.com/SladeThe/yav/internal"
)

var (
	requiredWithAnyIntFuncs map[string]yav.ValidateFunc[int]
)

func RequiredWithAnyInt() accumulators.RequiredWithAny[int] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyInt)
}

func provideRequiredWithAnyInt(names string, required bool) yav.ValidateFunc[int] {
	if !required {
		return OmitEmpty[int]
	}

	if validateFunc, ok := requiredWithAnyIntFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyIntFuncs, names, requiredWithAnyInt(names))
}

func requiredWithAnyInt(names string) yav.ValidateFunc[int] {
	return func(name string, value int) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyInt8Funcs map[string]yav.ValidateFunc[int8]
)

func RequiredWithAnyInt8() accumulators.RequiredWithAny[int8] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyInt8)
}

func provideRequiredWithAnyInt8(names string, required bool) yav.ValidateFunc[int8] {
	if !required {
		return OmitEmpty[int8]
	}

	if validateFunc, ok := requiredWithAnyInt8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyInt8Funcs, names, requiredWithAnyInt8(names))
}

func requiredWithAnyInt8(names string) yav.ValidateFunc[int8] {
	return func(name string, value int8) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyInt16Funcs map[string]yav.ValidateFunc[int16]
)

func RequiredWithAnyInt16() accumulators.RequiredWithAny[int16] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyInt16)
}

func provideRequiredWithAnyInt16(names string, required bool) yav.ValidateFunc[int16] {
	if !required {
		return OmitEmpty[int16]
	}

	if validateFunc, ok := requiredWithAnyInt16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyInt16Funcs, names, requiredWithAnyInt16(names))
}

func requiredWithAnyInt16(names string) yav.ValidateFunc[int16] {
	return func(name string, value int16) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyInt32Funcs map[string]yav.ValidateFunc[int32]
)

func RequiredWithAnyInt32() accumulators.RequiredWithAny[int32] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyInt32)
}

func provideRequiredWithAnyInt32(names string, required bool) yav.ValidateFunc[int32] {
	if !required {
		return OmitEmpty[int32]
	}

	if validateFunc, ok := requiredWithAnyInt32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyInt32Funcs, names, requiredWithAnyInt32(names))
}

func requiredWithAnyInt32(names string) yav.ValidateFunc[int32] {
	return func(name string, value int32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyInt64Funcs map[string]yav.ValidateFunc[int64]
)

func RequiredWithAnyInt64() accumulators.RequiredWithAny[int64] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyInt64)
}

func provideRequiredWithAnyInt64(names string, required bool) yav.ValidateFunc[int64] {
	if !required {
		return OmitEmpty[int64]
	}

	if validateFunc, ok := requiredWithAnyInt64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyInt64Funcs, names, requiredWithAnyInt64(names))
}

func requiredWithAnyInt64(names string) yav.ValidateFunc[int64] {
	return func(name string, value int64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyUintFuncs map[string]yav.ValidateFunc[uint]
)

func RequiredWithAnyUint() accumulators.RequiredWithAny[uint] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyUint)
}

func provideRequiredWithAnyUint(names string, required bool) yav.ValidateFunc[uint] {
	if !required {
		return OmitEmpty[uint]
	}

	if validateFunc, ok := requiredWithAnyUintFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyUintFuncs, names, requiredWithAnyUint(names))
}

func requiredWithAnyUint(names string) yav.ValidateFunc[uint] {
	return func(name string, value uint) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyUint8Funcs map[string]yav.ValidateFunc[uint8]
)

func RequiredWithAnyUint8() accumulators.RequiredWithAny[uint8] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyUint8)
}

func provideRequiredWithAnyUint8(names string, required bool) yav.ValidateFunc[uint8] {
	if !required {
		return OmitEmpty[uint8]
	}

	if validateFunc, ok := requiredWithAnyUint8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyUint8Funcs, names, requiredWithAnyUint8(names))
}

func requiredWithAnyUint8(names string) yav.ValidateFunc[uint8] {
	return func(name string, value uint8) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyUint16Funcs map[string]yav.ValidateFunc[uint16]
)

func RequiredWithAnyUint16() accumulators.RequiredWithAny[uint16] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyUint16)
}

func provideRequiredWithAnyUint16(names string, required bool) yav.ValidateFunc[uint16] {
	if !required {
		return OmitEmpty[uint16]
	}

	if validateFunc, ok := requiredWithAnyUint16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyUint16Funcs, names, requiredWithAnyUint16(names))
}

func requiredWithAnyUint16(names string) yav.ValidateFunc[uint16] {
	return func(name string, value uint16) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyUint32Funcs map[string]yav.ValidateFunc[uint32]
)

func RequiredWithAnyUint32() accumulators.RequiredWithAny[uint32] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyUint32)
}

func provideRequiredWithAnyUint32(names string, required bool) yav.ValidateFunc[uint32] {
	if !required {
		return OmitEmpty[uint32]
	}

	if validateFunc, ok := requiredWithAnyUint32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyUint32Funcs, names, requiredWithAnyUint32(names))
}

func requiredWithAnyUint32(names string) yav.ValidateFunc[uint32] {
	return func(name string, value uint32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyUint64Funcs map[string]yav.ValidateFunc[uint64]
)

func RequiredWithAnyUint64() accumulators.RequiredWithAny[uint64] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyUint64)
}

func provideRequiredWithAnyUint64(names string, required bool) yav.ValidateFunc[uint64] {
	if !required {
		return OmitEmpty[uint64]
	}

	if validateFunc, ok := requiredWithAnyUint64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyUint64Funcs, names, requiredWithAnyUint64(names))
}

func requiredWithAnyUint64(names string) yav.ValidateFunc[uint64] {
	return func(name string, value uint64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyFloat32Funcs map[string]yav.ValidateFunc[float32]
)

func RequiredWithAnyFloat32() accumulators.RequiredWithAny[float32] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyFloat32)
}

func provideRequiredWithAnyFloat32(names string, required bool) yav.ValidateFunc[float32] {
	if !required {
		return OmitEmpty[float32]
	}

	if validateFunc, ok := requiredWithAnyFloat32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyFloat32Funcs, names, requiredWithAnyFloat32(names))
}

func requiredWithAnyFloat32(names string) yav.ValidateFunc[float32] {
	return func(name string, value float32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAnyFloat64Funcs map[string]yav.ValidateFunc[float64]
)

func RequiredWithAnyFloat64() accumulators.RequiredWithAny[float64] {
	return accumulators.NewRequiredWithAny(provideRequiredWithAnyFloat64)
}

func provideRequiredWithAnyFloat64(names string, required bool) yav.ValidateFunc[float64] {
	if !required {
		return OmitEmpty[float64]
	}

	if validateFunc, ok := requiredWithAnyFloat64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAnyFloat64Funcs, names, requiredWithAnyFloat64(names))
}

func requiredWithAnyFloat64(names string) yav.ValidateFunc[float64] {
	return func(name string, value float64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
