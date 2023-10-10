package vstring

import (
	"regexp"

	"github.com/SladeThe/yav"
)

var (
	base64Regex       = regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$")
	base64URLRegex    = regexp.MustCompile("^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$")
	base64RawURLRegex = regexp.MustCompile("^(?:[A-Za-z0-9-_]{4})*[A-Za-z0-9-_]{2,4}$")
)

func Base64(name string, value string) (stop bool, err error) {
	if !base64Regex.MatchString(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Base64URL(name string, value string) (stop bool, err error) {
	if !base64URLRegex.MatchString(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64URL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Base64RawURL(name string, value string) (stop bool, err error) {
	if !base64RawURLRegex.MatchString(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64RawURL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
