// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package vnumber

import "github.com/SladeThe/yav"

func RequiredIfInt(conditionString string, condition bool) yav.ValidateFunc[int] {
	if !condition {
		return OmitEmpty[int]
	}

	return requiredIfInt(conditionString).validate
}

type requiredIfInt string

func (r requiredIfInt) validate(name string, value int) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfInt8(conditionString string, condition bool) yav.ValidateFunc[int8] {
	if !condition {
		return OmitEmpty[int8]
	}

	return requiredIfInt8(conditionString).validate
}

type requiredIfInt8 string

func (r requiredIfInt8) validate(name string, value int8) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfInt16(conditionString string, condition bool) yav.ValidateFunc[int16] {
	if !condition {
		return OmitEmpty[int16]
	}

	return requiredIfInt16(conditionString).validate
}

type requiredIfInt16 string

func (r requiredIfInt16) validate(name string, value int16) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfInt32(conditionString string, condition bool) yav.ValidateFunc[int32] {
	if !condition {
		return OmitEmpty[int32]
	}

	return requiredIfInt32(conditionString).validate
}

type requiredIfInt32 string

func (r requiredIfInt32) validate(name string, value int32) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfInt64(conditionString string, condition bool) yav.ValidateFunc[int64] {
	if !condition {
		return OmitEmpty[int64]
	}

	return requiredIfInt64(conditionString).validate
}

type requiredIfInt64 string

func (r requiredIfInt64) validate(name string, value int64) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfUint(conditionString string, condition bool) yav.ValidateFunc[uint] {
	if !condition {
		return OmitEmpty[uint]
	}

	return requiredIfUint(conditionString).validate
}

type requiredIfUint string

func (r requiredIfUint) validate(name string, value uint) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfUint8(conditionString string, condition bool) yav.ValidateFunc[uint8] {
	if !condition {
		return OmitEmpty[uint8]
	}

	return requiredIfUint8(conditionString).validate
}

type requiredIfUint8 string

func (r requiredIfUint8) validate(name string, value uint8) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfUint16(conditionString string, condition bool) yav.ValidateFunc[uint16] {
	if !condition {
		return OmitEmpty[uint16]
	}

	return requiredIfUint16(conditionString).validate
}

type requiredIfUint16 string

func (r requiredIfUint16) validate(name string, value uint16) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfUint32(conditionString string, condition bool) yav.ValidateFunc[uint32] {
	if !condition {
		return OmitEmpty[uint32]
	}

	return requiredIfUint32(conditionString).validate
}

type requiredIfUint32 string

func (r requiredIfUint32) validate(name string, value uint32) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfUint64(conditionString string, condition bool) yav.ValidateFunc[uint64] {
	if !condition {
		return OmitEmpty[uint64]
	}

	return requiredIfUint64(conditionString).validate
}

type requiredIfUint64 string

func (r requiredIfUint64) validate(name string, value uint64) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfFloat32(conditionString string, condition bool) yav.ValidateFunc[float32] {
	if !condition {
		return OmitEmpty[float32]
	}

	return requiredIfFloat32(conditionString).validate
}

type requiredIfFloat32 string

func (r requiredIfFloat32) validate(name string, value float32) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}

func RequiredIfFloat64(conditionString string, condition bool) yav.ValidateFunc[float64] {
	if !condition {
		return OmitEmpty[float64]
	}

	return requiredIfFloat64(conditionString).validate
}

type requiredIfFloat64 string

func (r requiredIfFloat64) validate(name string, value float64) (stop bool, err error) {
	if value == 0 {
		return true, yav.Error{
			CheckName: yav.CheckNameRequiredIf,
			Parameter: string(r),
			ValueName: name,
		}
	}

	return false, nil
}
