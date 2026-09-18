package vstring

import (
	"net/url"
	"strings"

	"github.com/SladeThe/yav"
)

func URI(name string, value string) (stop bool, err error) {
	if !isURX(value, true) {
		return true, yav.Error{
			CheckName: yav.CheckNameURI,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func URL(name string, value string) (stop bool, err error) {
	if !isURX(value, false) {
		return true, yav.Error{
			CheckName: yav.CheckNameURL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func isURX(value string, isURI bool) bool {
	if i := strings.IndexRune(value, '#'); i >= 0 {
		if !isURXComponent(value[i+1:]) {
			return false
		}

		value = value[:i]
	}

	if value == "" {
		return false
	}

	if isSimpleHTTPURL(value) {
		return true
	}

	u, err := url.ParseRequestURI(value)
	return err == nil && (isURI || u.Scheme != "")
}

func isSimpleHTTPURL(value string) bool {
	switch {
	case strings.HasPrefix(value, "https://"):
		value = value[len("https://"):]
	case strings.HasPrefix(value, "http://"):
		value = value[len("http://"):]
	default:
		return false
	}

	index := 0

	for index < len(value) {
		c := value[index]

		if lower := c | ('a' - 'A'); (lower < 'a' || lower > 'z') && (c < '0' || c > '9') && c != '.' && c != '-' {
			break
		}

		index++
	}

	if index == 0 {
		return false
	}

	if index < len(value) && value[index] == ':' {
		index++

		for index < len(value) && '0' <= value[index] && value[index] <= '9' {
			index++
		}
	}

	return index == len(value) ||
		(value[index] == '/' || value[index] == '?') && isURXComponent(value[index:])
}

// Equivalent regexp: `^(?:[^\x00-\x1f\x7f%]|%[0-9a-fA-F]{2})*$`
func isURXComponent(value string) bool {
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c < ' ' || c == '\x7f' {
			return false
		}

		if c == '%' {
			if i+2 >= len(value) || !isHexDigit(value[i+1]) || !isHexDigit(value[i+2]) {
				return false
			}

			i += 2
		}
	}

	return true
}

func isHexDigit(c byte) bool {
	return c >= '0' && c <= '9' ||
		c >= 'a' && c <= 'f' ||
		c >= 'A' && c <= 'F'
}
