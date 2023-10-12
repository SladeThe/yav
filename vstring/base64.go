package vstring

import (
	"github.com/SladeThe/yav"
)

var (
	base64Charset = func() [256]bool {
		var chars [256]bool

		for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/" {
			chars[r] = true
		}

		return chars
	}()

	base64URLCharset = func() [256]bool {
		var chars [256]bool

		for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_" {
			chars[r] = true
		}

		return chars
	}()
)

func Base64(name string, value string) (stop bool, err error) {
	if !isBase64(value, &base64Charset) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Base64Raw(name string, value string) (stop bool, err error) {
	if !isBase64Raw(value, &base64Charset) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64Raw,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Base64URL(name string, value string) (stop bool, err error) {
	if !isBase64(value, &base64URLCharset) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64URL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func Base64RawURL(name string, value string) (stop bool, err error) {
	if !isBase64Raw(value, &base64URLCharset) {
		return true, yav.Error{
			CheckName: yav.CheckNameBase64RawURL,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func isBase64(value string, charset *[256]bool) bool {
	if value == "" {
		return true
	}

	if len(value)%4 != 0 {
		return false
	}

	paddingLength := 0

	if value[len(value)-1] == '=' {
		paddingLength++

		if value[len(value)-2] == '=' {
			paddingLength++
		}
	}

	for _, r := range []byte(value[:len(value)-paddingLength]) {
		if !charset[r] {
			return false
		}
	}

	return true
}

func isBase64Raw(value string, charset *[256]bool) bool {
	if value == "" {
		return true
	}

	if len(value)%4 == 1 {
		return false
	}

	for _, r := range []byte(value) {
		if !charset[r] {
			return false
		}
	}

	return true
}
