package vstring

import (
	"strings"
	"unicode/utf8"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/common"
)

func Email(name string, value string) (stop bool, err error) {
	if !isEmail(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameEmail,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

// Equivalent regexp:
// `^(?:[a-zA-Z0-9!#$%&'*+/=?^_\x60{|}~\-\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}]+(?:\.[a-zA-Z0-9!#$%&'*+/=?^_\x60{|}~\-\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}]+)*|"(?:[\x09\x20\x21\x23-\x5b\x5d-\x7e\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}]|\\[\x09\x20-\x7e\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}])+")@(?:[a-zA-Z0-9\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}](?:[a-zA-Z0-9\-\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}]*[a-zA-Z0-9\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}])?\.)+[a-zA-Z\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}](?:[a-zA-Z0-9\-\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}]*[a-zA-Z0-9\x{00A0}-\x{D7FF}\x{F900}-\x{FDCF}\x{FDF0}-\x{FFEF}])?$`
func isEmail(value string) bool {
	atIndex := strings.LastIndexByte(value, '@')
	return atIndex > 0 && isEmailLocal(value[:atIndex]) && isEmailDomain(value[atIndex+1:])
}

func isEmailLocal(value string) bool {
	if value[0] == '"' {
		return isEmailQuotedLocal(value)
	}

	previousDot := true
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c == '.' {
			if previousDot {
				return false
			}

			previousDot = true
			continue
		}

		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(value[i:])
			if !isEmailUnicode(r) {
				return false
			}

			i += size - 1
		} else if !common.IsAlphanumeric(rune(c)) {
			switch c {
			case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '/', '=', '?', '^', '_', '`', '{', '|', '}', '~':
			default:
				return false
			}
		}

		previousDot = false
	}

	return !previousDot
}

func isEmailQuotedLocal(value string) bool {
	lastIndex := len(value) - 1
	if lastIndex < 2 || value[lastIndex] != '"' {
		return false
	}

	for i := 1; i < lastIndex; i++ {
		c := value[i]
		if c == '\\' {
			i++
			if i == lastIndex {
				return false
			}

			c = value[i]
		} else if c == '"' {
			return false
		}

		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(value[i:])
			if !isEmailUnicode(r) {
				return false
			}

			i += size - 1
		} else if c != '\t' && (c < ' ' || c > '~') {
			return false
		}
	}

	return true
}

func isEmailDomain(value string) bool {
	if value == "" {
		return false
	}

	labelStart := 0
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c == '.' {
			if i == labelStart || value[i-1] == '-' {
				return false
			}

			labelStart = i + 1
			continue
		}

		if c == '-' {
			if i == labelStart {
				return false
			}

			continue
		}

		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(value[i:])
			if !isEmailUnicode(r) {
				return false
			}

			i += size - 1
		} else if !common.IsAlphanumeric(rune(c)) {
			return false
		}
	}

	// The scan already validates each label's first character.
	return labelStart > 0 && labelStart < len(value) && value[len(value)-1] != '-' &&
		!common.IsDigit(rune(value[labelStart]))
}

// Unicode ranges accepted by go-playground's email regexp.
func isEmailUnicode(r rune) bool {
	return r >= 0x00A0 && r <= 0xD7FF ||
		r >= 0xF900 && r <= 0xFDCF ||
		r >= 0xFDF0 && r <= 0xFFEF
}
