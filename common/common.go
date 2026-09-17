package common

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	CharactersSpecial = ",!@#$%^&*)(+=._-"
)

func IsLowerAlpha(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsUpperAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsAlpha(r rune) bool {
	return IsLowerAlpha(r) || IsUpperAlpha(r)
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsAlphanumeric(r rune) bool {
	return IsAlpha(r) || IsDigit(r)
}

func IsSpecialCharacter(r rune) bool {
	return strings.ContainsRune(CharactersSpecial, r)
}

// IsText checks for valid UTF-8 containing only graphic characters and CR/LF line breaks.
// No ring bells and other control characters. The value is allowed to be empty.
func IsText(value string) bool {
	for index, r := range value {
		if r == utf8.RuneError {
			if _, width := utf8.DecodeRuneInString(value[index:]); width == 1 {
				return false
			}
		}

		if r != '\r' && r != '\n' && !unicode.IsGraphic(r) {
			return false
		}
	}

	return true
}

// IsTitle checks for a non-empty, valid UTF-8 line with neither leading nor trailing spaces.
// It must contain only printable characters and single ASCII spaces (a space must not follow a space).
// No tabulations, no unbreakable spaces, no ring bells and other control characters.
func IsTitle(value string) bool {
	if value == "" || len(value) != len(strings.TrimSpace(value)) {
		return false
	}

	lastSpace := false

	for index, r := range value {
		if r == utf8.RuneError {
			if _, width := utf8.DecodeRuneInString(value[index:]); width == 1 {
				return false
			}
		}

		if !unicode.IsPrint(r) {
			return false
		}

		if r == ' ' && lastSpace {
			return false
		}

		lastSpace = r == ' '
	}

	return true
}
