package vtime

import (
	"time"

	"github.com/SladeThe/yav"
)

func OmitEmpty(_ string, value time.Time) (stop bool, err error) {
	return value.IsZero(), nil
}

func Required(name string, value time.Time) (stop bool, err error) {
	if value.IsZero() {
		return false, yav.Error{
			CheckName: yav.CheckNameRequired,
			ValueName: name,
		}
	}

	return false, nil
}
