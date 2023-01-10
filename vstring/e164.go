package vstring

import (
	"regexp"

	"github.com/SladeThe/yav"
)

var (
	e164Regex = regexp.MustCompile("^\\+[1-9]?[0-9]{7,14}$")
)

func E164(name string, value string) (stop bool, err error) {
	if !e164Regex.MatchString(value) {
		return false, yav.Error{
			CheckName: yav.CheckNameE164,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
