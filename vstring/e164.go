package vstring

import (
	"github.com/SladeThe/yav"
)

const (
	minE164Length = 8
	maxE164Length = 16
)

func E164(name string, value string) (stop bool, err error) {
	if !isE164(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameE164,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func isE164(value string) bool {
	if len(value) < minE164Length || len(value) > maxE164Length {
		return false
	}

	if value[0] != '+' || (value[1] < '1' || value[1] > '9') {
		return false
	}

	for _, r := range []byte(value[2:]) {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}
