package vstring

import (
	"regexp"

	"github.com/SladeThe/yav"
)

var (
	uuidRegex = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
)

func IsUUID(name string, value string) (stop bool, err error) {
	if !uuidRegex.MatchString(value) {
		return false, yav.Error{
			CheckName: yav.CheckNameUUID,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
