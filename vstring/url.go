package vstring

import (
	"net/url"
	"strings"

	"github.com/SladeThe/yav"
)

func IsURI(name string, value string) (stop bool, err error) {
	if !isURX(value, true) {
		return false, yav.Error{
			CheckName: yav.CheckNameURI,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func IsURL(name string, value string) (stop bool, err error) {
	if !isURX(value, false) {
		return false, yav.Error{
			CheckName: yav.CheckNameURL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func isURX(value string, isURI bool) bool {
	if i := strings.IndexRune(value, '#'); i >= 0 {
		value = value[:i]
	}

	if value == "" {
		return false
	}

	u, err := url.ParseRequestURI(value)
	return err == nil && (isURI || u.Scheme != "")
}
