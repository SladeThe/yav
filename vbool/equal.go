package vbool

import (
	"github.com/SladeThe/yav"
)

func Equal(parameter bool) yav.ValidateFunc[bool] {
	if parameter {
		return equalTrue
	}

	return equalFalse
}

func NotEqual(parameter bool) yav.ValidateFunc[bool] {
	if parameter {
		return notEqualTrue
	}

	return notEqualFalse
}

func equalTrue(name string, value bool) (stop bool, err error) {
	if !value {
		return true, yav.Error{
			CheckName: yav.CheckNameEqual,
			Parameter: "true",
			ValueName: name,
		}
	}

	return false, nil
}

func equalFalse(name string, value bool) (stop bool, err error) {
	if value {
		return true, yav.Error{
			CheckName: yav.CheckNameEqual,
			Parameter: "false",
			ValueName: name,
		}
	}

	return false, nil
}

func notEqualTrue(name string, value bool) (stop bool, err error) {
	if value {
		return true, yav.Error{
			CheckName: yav.CheckNameNotEqual,
			Parameter: "true",
			ValueName: name,
		}
	}

	return false, nil
}

func notEqualFalse(name string, value bool) (stop bool, err error) {
	if !value {
		return true, yav.Error{
			CheckName: yav.CheckNameNotEqual,
			Parameter: "false",
			ValueName: name,
		}
	}

	return false, nil
}
