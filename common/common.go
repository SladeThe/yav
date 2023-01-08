package common

import (
	"strings"
	"unicode"
)

const (
	CharactersLowerAlpha = "abcdefghijklmnopqrstuvwxyz"
	CharactersUpperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharactersAlpha      = CharactersLowerAlpha + CharactersUpperAlpha
	CharactersDigit      = "0123456789"
	CharactersSpecial    = ",!@#$%^&*)(+=._-"
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

// IsText checks the given value to contain only visible characters and various whitespaces, including line breaks.
// No ring bells and other control characters.
func IsText(value string) bool {
	for _, r := range value {
		if r != '\r' && r != '\n' && !unicode.IsGraphic(r) {
			return false
		}
	}

	return true
}

// IsTitle checks the given value is a single line string with neither leading nor trailing spaces.
// It must contain only printable characters and single ASCII spaces (a space must not follow a space).
// No tabulations, no unbreakable spaces, no ring bells and other control characters.
func IsTitle(value string) bool {
	// TODO improve performance

	if len(value) != len(strings.TrimSpace(value)) || strings.Contains(value, "  ") {
		return false
	}

	for _, r := range value {
		if !unicode.IsPrint(r) {
			return false
		}
	}

	return true
}
