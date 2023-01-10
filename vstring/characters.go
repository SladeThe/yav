package vstring

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/common"
)

func Alpha(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if !common.IsAlpha(r) {
			return false, yav.Error{
				CheckName: yav.CheckNameAlpha,
				ValueName: name,
				Value:     value,
			}
		}
	}

	return false, nil
}

func Alphanumeric(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if !common.IsAlphanumeric(r) {
			return false, yav.Error{
				CheckName: yav.CheckNameAlphanumeric,
				ValueName: name,
				Value:     value,
			}
		}
	}

	return false, nil
}

func Lowercase(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if unicode.IsUpper(r) {
			return false, yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: name,
				Value:     value,
			}
		}
	}

	return false, nil
}

func Uppercase(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if unicode.IsLower(r) {
			return false, yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: name,
				Value:     value,
			}
		}
	}

	return false, nil
}

func ContainsAlpha(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if common.IsAlpha(r) {
			return false, nil
		}
	}

	return false, yav.Error{
		CheckName: yav.CheckNameContainsAlpha,
		ValueName: name,
		Value:     value,
	}
}

func ContainsLowerAlpha(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if common.IsLowerAlpha(r) {
			return false, nil
		}
	}

	return false, yav.Error{
		CheckName: yav.CheckNameContainsLowerAlpha,
		ValueName: name,
		Value:     value,
	}
}

func ContainsUpperAlpha(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if common.IsUpperAlpha(r) {
			return false, nil
		}
	}

	return false, yav.Error{
		CheckName: yav.CheckNameContainsUpperAlpha,
		ValueName: name,
		Value:     value,
	}
}

func ContainsDigit(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if common.IsDigit(r) {
			return false, nil
		}
	}

	return false, yav.Error{
		CheckName: yav.CheckNameContainsDigit,
		ValueName: name,
		Value:     value,
	}
}

func ContainsSpecialCharacter(name string, value string) (stop bool, err error) {
	if strings.ContainsAny(value, common.CharactersSpecial) {
		return false, nil
	}

	return false, yav.Error{
		CheckName: yav.CheckNameContainsSpecialCharacter,
		ValueName: name,
		Value:     value,
	}
}

func ExcludesWhitespace(name string, value string) (stop bool, err error) {
	for _, r := range value {
		if unicode.IsSpace(r) {
			return false, yav.Error{
				CheckName: yav.CheckNameExcludesWhitespace,
				ValueName: name,
				Value:     value,
			}
		}
	}

	return false, nil
}

func StartsWithAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeRuneInString(value); !common.IsAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameStartsWithAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func StartsWithLowerAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeRuneInString(value); !common.IsLowerAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameStartsWithLowerAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func StartsWithUpperAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeRuneInString(value); !common.IsUpperAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameStartsWithUpperAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func StartsWithDigit(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeRuneInString(value); !common.IsDigit(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameStartsWithDigit,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func StartsWithSpecialCharacter(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeRuneInString(value); !common.IsSpecialCharacter(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameStartsWithSpecialCharacter,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func EndsWithAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeLastRuneInString(value); !common.IsAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameEndsWithAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func EndsWithLowerAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeLastRuneInString(value); !common.IsLowerAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameEndsWithLowerAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func EndsWithUpperAlpha(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeLastRuneInString(value); !common.IsUpperAlpha(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameEndsWithUpperAlpha,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func EndsWithDigit(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeLastRuneInString(value); !common.IsDigit(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameEndsWithDigit,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func EndsWithSpecialCharacter(name string, value string) (stop bool, err error) {
	if r, _ := utf8.DecodeLastRuneInString(value); !common.IsSpecialCharacter(r) {
		return false, yav.Error{
			CheckName: yav.CheckNameEndsWithSpecialCharacter,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

// Text validates string value using common.IsText.
func Text(name string, value string) (stop bool, err error) {
	if common.IsText(value) {
		return false, nil
	}

	return false, yav.Error{
		CheckName: yav.CheckNameText,
		ValueName: name,
		Value:     value,
	}
}

// Title validates string value using common.IsTitle.
func Title(name string, value string) (stop bool, err error) {
	if common.IsTitle(value) {
		return false, nil
	}

	return false, yav.Error{
		CheckName: yav.CheckNameTitle,
		ValueName: name,
		Value:     value,
	}
}
