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
	excludedWithAnyIntFuncs map[string]yav.ValidateFunc[int]
)

func ExcludedWithAnyInt() accumulators.ExcludedWithAny[int] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyInt)
}

func provideExcludedWithAnyInt(names string, excluded bool) yav.ValidateFunc[int] {
	if !excluded {
		return yav.Next[int]
	}

	if validateFunc, ok := excludedWithAnyIntFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyIntFuncs, names, excludedWithAnyInt(names))
}

func excludedWithAnyInt(names string) yav.ValidateFunc[int] {
	return func(name string, value int) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyInt8Funcs map[string]yav.ValidateFunc[int8]
)

func ExcludedWithAnyInt8() accumulators.ExcludedWithAny[int8] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyInt8)
}

func provideExcludedWithAnyInt8(names string, excluded bool) yav.ValidateFunc[int8] {
	if !excluded {
		return yav.Next[int8]
	}

	if validateFunc, ok := excludedWithAnyInt8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyInt8Funcs, names, excludedWithAnyInt8(names))
}

func excludedWithAnyInt8(names string) yav.ValidateFunc[int8] {
	return func(name string, value int8) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyInt16Funcs map[string]yav.ValidateFunc[int16]
)

func ExcludedWithAnyInt16() accumulators.ExcludedWithAny[int16] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyInt16)
}

func provideExcludedWithAnyInt16(names string, excluded bool) yav.ValidateFunc[int16] {
	if !excluded {
		return yav.Next[int16]
	}

	if validateFunc, ok := excludedWithAnyInt16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyInt16Funcs, names, excludedWithAnyInt16(names))
}

func excludedWithAnyInt16(names string) yav.ValidateFunc[int16] {
	return func(name string, value int16) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyInt32Funcs map[string]yav.ValidateFunc[int32]
)

func ExcludedWithAnyInt32() accumulators.ExcludedWithAny[int32] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyInt32)
}

func provideExcludedWithAnyInt32(names string, excluded bool) yav.ValidateFunc[int32] {
	if !excluded {
		return yav.Next[int32]
	}

	if validateFunc, ok := excludedWithAnyInt32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyInt32Funcs, names, excludedWithAnyInt32(names))
}

func excludedWithAnyInt32(names string) yav.ValidateFunc[int32] {
	return func(name string, value int32) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyInt64Funcs map[string]yav.ValidateFunc[int64]
)

func ExcludedWithAnyInt64() accumulators.ExcludedWithAny[int64] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyInt64)
}

func provideExcludedWithAnyInt64(names string, excluded bool) yav.ValidateFunc[int64] {
	if !excluded {
		return yav.Next[int64]
	}

	if validateFunc, ok := excludedWithAnyInt64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyInt64Funcs, names, excludedWithAnyInt64(names))
}

func excludedWithAnyInt64(names string) yav.ValidateFunc[int64] {
	return func(name string, value int64) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyUintFuncs map[string]yav.ValidateFunc[uint]
)

func ExcludedWithAnyUint() accumulators.ExcludedWithAny[uint] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyUint)
}

func provideExcludedWithAnyUint(names string, excluded bool) yav.ValidateFunc[uint] {
	if !excluded {
		return yav.Next[uint]
	}

	if validateFunc, ok := excludedWithAnyUintFuncs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyUintFuncs, names, excludedWithAnyUint(names))
}

func excludedWithAnyUint(names string) yav.ValidateFunc[uint] {
	return func(name string, value uint) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyUint8Funcs map[string]yav.ValidateFunc[uint8]
)

func ExcludedWithAnyUint8() accumulators.ExcludedWithAny[uint8] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyUint8)
}

func provideExcludedWithAnyUint8(names string, excluded bool) yav.ValidateFunc[uint8] {
	if !excluded {
		return yav.Next[uint8]
	}

	if validateFunc, ok := excludedWithAnyUint8Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyUint8Funcs, names, excludedWithAnyUint8(names))
}

func excludedWithAnyUint8(names string) yav.ValidateFunc[uint8] {
	return func(name string, value uint8) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyUint16Funcs map[string]yav.ValidateFunc[uint16]
)

func ExcludedWithAnyUint16() accumulators.ExcludedWithAny[uint16] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyUint16)
}

func provideExcludedWithAnyUint16(names string, excluded bool) yav.ValidateFunc[uint16] {
	if !excluded {
		return yav.Next[uint16]
	}

	if validateFunc, ok := excludedWithAnyUint16Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyUint16Funcs, names, excludedWithAnyUint16(names))
}

func excludedWithAnyUint16(names string) yav.ValidateFunc[uint16] {
	return func(name string, value uint16) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyUint32Funcs map[string]yav.ValidateFunc[uint32]
)

func ExcludedWithAnyUint32() accumulators.ExcludedWithAny[uint32] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyUint32)
}

func provideExcludedWithAnyUint32(names string, excluded bool) yav.ValidateFunc[uint32] {
	if !excluded {
		return yav.Next[uint32]
	}

	if validateFunc, ok := excludedWithAnyUint32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyUint32Funcs, names, excludedWithAnyUint32(names))
}

func excludedWithAnyUint32(names string) yav.ValidateFunc[uint32] {
	return func(name string, value uint32) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyUint64Funcs map[string]yav.ValidateFunc[uint64]
)

func ExcludedWithAnyUint64() accumulators.ExcludedWithAny[uint64] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyUint64)
}

func provideExcludedWithAnyUint64(names string, excluded bool) yav.ValidateFunc[uint64] {
	if !excluded {
		return yav.Next[uint64]
	}

	if validateFunc, ok := excludedWithAnyUint64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyUint64Funcs, names, excludedWithAnyUint64(names))
}

func excludedWithAnyUint64(names string) yav.ValidateFunc[uint64] {
	return func(name string, value uint64) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyFloat32Funcs map[string]yav.ValidateFunc[float32]
)

func ExcludedWithAnyFloat32() accumulators.ExcludedWithAny[float32] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyFloat32)
}

func provideExcludedWithAnyFloat32(names string, excluded bool) yav.ValidateFunc[float32] {
	if !excluded {
		return yav.Next[float32]
	}

	if validateFunc, ok := excludedWithAnyFloat32Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyFloat32Funcs, names, excludedWithAnyFloat32(names))
}

func excludedWithAnyFloat32(names string) yav.ValidateFunc[float32] {
	return func(name string, value float32) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}

var (
	excludedWithAnyFloat64Funcs map[string]yav.ValidateFunc[float64]
)

func ExcludedWithAnyFloat64() accumulators.ExcludedWithAny[float64] {
	return accumulators.NewExcludedWithAny(provideExcludedWithAnyFloat64)
}

func provideExcludedWithAnyFloat64(names string, excluded bool) yav.ValidateFunc[float64] {
	if !excluded {
		return yav.Next[float64]
	}

	if validateFunc, ok := excludedWithAnyFloat64Funcs[names]; ok {
		return validateFunc
	}

	return internal.RegisterMapEntry(&excludedWithAnyFloat64Funcs, names, excludedWithAnyFloat64(names))
}

func excludedWithAnyFloat64(names string) yav.ValidateFunc[float64] {
	return func(name string, value float64) (stop bool, err error) {
		if value != 0 {
			return true, yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: names,
				ValueName: name,
			}
		}

		return false, nil
	}
}
