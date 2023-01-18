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
	requiredWithAllIntFuncs map[string]yav.ValidateFunc[int]
)

func RequiredWithAllInt() accumulators.RequiredWithAll[int] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllInt)
}

func provideRequiredWithAllInt(names string, required bool) yav.ValidateFunc[int] {
	if !required {
		return OmitEmpty[int]
	}

	if validateFunc, ok := requiredWithAllIntFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllIntFuncs, names, requiredWithAllInt(names))
}

func requiredWithAllInt(names string) yav.ValidateFunc[int] {
	return func(name string, value int) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllInt8Funcs map[string]yav.ValidateFunc[int8]
)

func RequiredWithAllInt8() accumulators.RequiredWithAll[int8] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllInt8)
}

func provideRequiredWithAllInt8(names string, required bool) yav.ValidateFunc[int8] {
	if !required {
		return OmitEmpty[int8]
	}

	if validateFunc, ok := requiredWithAllInt8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllInt8Funcs, names, requiredWithAllInt8(names))
}

func requiredWithAllInt8(names string) yav.ValidateFunc[int8] {
	return func(name string, value int8) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllInt16Funcs map[string]yav.ValidateFunc[int16]
)

func RequiredWithAllInt16() accumulators.RequiredWithAll[int16] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllInt16)
}

func provideRequiredWithAllInt16(names string, required bool) yav.ValidateFunc[int16] {
	if !required {
		return OmitEmpty[int16]
	}

	if validateFunc, ok := requiredWithAllInt16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllInt16Funcs, names, requiredWithAllInt16(names))
}

func requiredWithAllInt16(names string) yav.ValidateFunc[int16] {
	return func(name string, value int16) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllInt32Funcs map[string]yav.ValidateFunc[int32]
)

func RequiredWithAllInt32() accumulators.RequiredWithAll[int32] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllInt32)
}

func provideRequiredWithAllInt32(names string, required bool) yav.ValidateFunc[int32] {
	if !required {
		return OmitEmpty[int32]
	}

	if validateFunc, ok := requiredWithAllInt32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllInt32Funcs, names, requiredWithAllInt32(names))
}

func requiredWithAllInt32(names string) yav.ValidateFunc[int32] {
	return func(name string, value int32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllInt64Funcs map[string]yav.ValidateFunc[int64]
)

func RequiredWithAllInt64() accumulators.RequiredWithAll[int64] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllInt64)
}

func provideRequiredWithAllInt64(names string, required bool) yav.ValidateFunc[int64] {
	if !required {
		return OmitEmpty[int64]
	}

	if validateFunc, ok := requiredWithAllInt64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllInt64Funcs, names, requiredWithAllInt64(names))
}

func requiredWithAllInt64(names string) yav.ValidateFunc[int64] {
	return func(name string, value int64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllUintFuncs map[string]yav.ValidateFunc[uint]
)

func RequiredWithAllUint() accumulators.RequiredWithAll[uint] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllUint)
}

func provideRequiredWithAllUint(names string, required bool) yav.ValidateFunc[uint] {
	if !required {
		return OmitEmpty[uint]
	}

	if validateFunc, ok := requiredWithAllUintFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllUintFuncs, names, requiredWithAllUint(names))
}

func requiredWithAllUint(names string) yav.ValidateFunc[uint] {
	return func(name string, value uint) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllUint8Funcs map[string]yav.ValidateFunc[uint8]
)

func RequiredWithAllUint8() accumulators.RequiredWithAll[uint8] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllUint8)
}

func provideRequiredWithAllUint8(names string, required bool) yav.ValidateFunc[uint8] {
	if !required {
		return OmitEmpty[uint8]
	}

	if validateFunc, ok := requiredWithAllUint8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllUint8Funcs, names, requiredWithAllUint8(names))
}

func requiredWithAllUint8(names string) yav.ValidateFunc[uint8] {
	return func(name string, value uint8) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllUint16Funcs map[string]yav.ValidateFunc[uint16]
)

func RequiredWithAllUint16() accumulators.RequiredWithAll[uint16] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllUint16)
}

func provideRequiredWithAllUint16(names string, required bool) yav.ValidateFunc[uint16] {
	if !required {
		return OmitEmpty[uint16]
	}

	if validateFunc, ok := requiredWithAllUint16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllUint16Funcs, names, requiredWithAllUint16(names))
}

func requiredWithAllUint16(names string) yav.ValidateFunc[uint16] {
	return func(name string, value uint16) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllUint32Funcs map[string]yav.ValidateFunc[uint32]
)

func RequiredWithAllUint32() accumulators.RequiredWithAll[uint32] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllUint32)
}

func provideRequiredWithAllUint32(names string, required bool) yav.ValidateFunc[uint32] {
	if !required {
		return OmitEmpty[uint32]
	}

	if validateFunc, ok := requiredWithAllUint32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllUint32Funcs, names, requiredWithAllUint32(names))
}

func requiredWithAllUint32(names string) yav.ValidateFunc[uint32] {
	return func(name string, value uint32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllUint64Funcs map[string]yav.ValidateFunc[uint64]
)

func RequiredWithAllUint64() accumulators.RequiredWithAll[uint64] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllUint64)
}

func provideRequiredWithAllUint64(names string, required bool) yav.ValidateFunc[uint64] {
	if !required {
		return OmitEmpty[uint64]
	}

	if validateFunc, ok := requiredWithAllUint64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllUint64Funcs, names, requiredWithAllUint64(names))
}

func requiredWithAllUint64(names string) yav.ValidateFunc[uint64] {
	return func(name string, value uint64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllFloat32Funcs map[string]yav.ValidateFunc[float32]
)

func RequiredWithAllFloat32() accumulators.RequiredWithAll[float32] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllFloat32)
}

func provideRequiredWithAllFloat32(names string, required bool) yav.ValidateFunc[float32] {
	if !required {
		return OmitEmpty[float32]
	}

	if validateFunc, ok := requiredWithAllFloat32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllFloat32Funcs, names, requiredWithAllFloat32(names))
}

func requiredWithAllFloat32(names string) yav.ValidateFunc[float32] {
	return func(name string, value float32) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	requiredWithAllFloat64Funcs map[string]yav.ValidateFunc[float64]
)

func RequiredWithAllFloat64() accumulators.RequiredWithAll[float64] {
	return accumulators.NewRequiredWithAll(provideRequiredWithAllFloat64)
}

func provideRequiredWithAllFloat64(names string, required bool) yav.ValidateFunc[float64] {
	if !required {
		return OmitEmpty[float64]
	}

	if validateFunc, ok := requiredWithAllFloat64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&requiredWithAllFloat64Funcs, names, requiredWithAllFloat64(names))
}

func requiredWithAllFloat64(names string) yav.ValidateFunc[float64] {
	return func(name string, value float64) (stop bool, err error) {
		if value == 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
