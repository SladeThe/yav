package yav

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

const (
	CheckNameRequired           = "required"
	CheckNameRequiredWithAny    = "required_with"
	CheckNameRequiredWithoutAny = "required_without"
	CheckNameRequiredWithAll    = "required_with_all"
	CheckNameRequiredWithoutAll = "required_without_all"

	CheckNameMin    = "min"
	CheckNameMax    = "max"
	CheckNameEmail  = "email"
	CheckNameE164   = "e164"
	CheckNameUUID   = "uuid"
	CheckNameEqual  = "eq"
	CheckNameOneOf  = "oneof"
	CheckNameUnique = "unique"

	CheckNameAlpha        = "alpha"
	CheckNameAlphanumeric = "alphanum"
	CheckNameLowercase    = "lowercase"
	CheckNameUppercase    = "uppercase"

	CheckNameContainsAlpha            = "contains_alpha"
	CheckNameContainsLowerAlpha       = "contains_lower_alpha"
	CheckNameContainsUpperAlpha       = "contains_upper_alpha"
	CheckNameContainsDigit            = "contains_digit"
	CheckNameContainsSpecialCharacter = "contains_special_character"

	CheckNameExcludesWhitespace = "excludes_whitespace"

	CheckNameStartsWithAlpha            = "starts_with_alpha"
	CheckNameStartsWithLowerAlpha       = "starts_with_lower_alpha"
	CheckNameStartsWithUpperAlpha       = "starts_with_upper_alpha"
	CheckNameStartsWithDigit            = "starts_with_digit"
	CheckNameStartsWithSpecialCharacter = "starts_with_special_character"

	CheckNameEndsWithAlpha            = "ends_with_alpha"
	CheckNameEndsWithLowerAlpha       = "ends_with_lower_alpha"
	CheckNameEndsWithUpperAlpha       = "ends_with_upper_alpha"
	CheckNameEndsWithDigit            = "ends_with_digit"
	CheckNameEndsWithSpecialCharacter = "ends_with_special_character"

	CheckNameText  = "text"
	CheckNameTitle = "title"
)

type Error struct {
	CheckName string
	Parameter string

	ValueName string
	Value     any
}

func IsError(err error) bool {
	var validationErr Error
	return errors.As(err, &validationErr)
}

func (err Error) Is(target error) bool {
	return err == target
}

func (err Error) Error() string {
	if err.CheckName != "" && err.ValueName != "" {
		if err.Parameter != "" {
			return fmt.Sprintf("'%s' failed the '%s=%s' check", err.ValueName, err.CheckName, err.Parameter)
		}

		return fmt.Sprintf("'%s' failed the '%s' check", err.ValueName, err.CheckName)
	}

	if err.CheckName != "" {
		return fmt.Sprintf("validation failed: %s", err.CheckName)
	}

	return "unknown error"
}

func (err Error) WithParameter(parameter string) Error {
	err.Parameter = parameter
	return err
}

func (err Error) WithValueName(name string) Error {
	err.ValueName = name
	return err
}

func (err Error) WithValue(value any) Error {
	err.Value = value
	return err
}

func (err Error) WithNamedValue(name string, value any) Error {
	err.ValueName = name
	err.Value = value
	return err
}

func (err Error) ValueAsString() string {
	switch v := err.Value.(type) {
	case string:
		return v
	case *string:
		if v != nil {
			return *v
		}
	case fmt.Stringer:
		return v.String()
	}

	return ""
}

func WithNamespace(namespace string, err error) error {
	if err == nil {
		return nil
	}

	if validationErr, ok := err.(Error); ok {
		if validationErr.ValueName == "" {
			return err
		}

		validationErr.ValueName = namespace + "." + validationErr.ValueName
		return validationErr
	}

	partialErrs := multierr.Errors(err) // TODO update in-place ?
	if len(partialErrs) <= 1 {
		return err
	}

	var combinedErr error

	for _, partialErr := range partialErrs {
		multierr.AppendInto(&combinedErr, WithNamespace(namespace, partialErr))
	}

	return combinedErr
}
