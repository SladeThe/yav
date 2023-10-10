package yav

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

const (
	CheckNameRequired           = "required"
	CheckNameRequiredIf         = "required_if"
	CheckNameRequiredUnless     = "required_unless"
	CheckNameRequiredWithAny    = "required_with"
	CheckNameRequiredWithoutAny = "required_without"
	CheckNameRequiredWithAll    = "required_with_all"
	CheckNameRequiredWithoutAll = "required_without_all"

	CheckNameExcludedIf         = "excluded_if"
	CheckNameExcludedUnless     = "excluded_unless"
	CheckNameExcludedWithAny    = "excluded_with"
	CheckNameExcludedWithoutAny = "excluded_without"
	CheckNameExcludedWithAll    = "excluded_with_all"
	CheckNameExcludedWithoutAll = "excluded_without_all"

	CheckNameMin                = "min"
	CheckNameMax                = "max"
	CheckNameGreaterThan        = "gt"
	CheckNameGreaterThanOrEqual = "gte"
	CheckNameLessThan           = "lt"
	CheckNameLessThanOrEqual    = "lte"

	CheckNameGreaterThanNamed        = "gtfield"
	CheckNameGreaterThanOrEqualNamed = "gtefield"
	CheckNameLessThanNamed           = "ltfield"
	CheckNameLessThanOrEqualNamed    = "ltefield"

	CheckNameEqual    = "eq"
	CheckNameNotEqual = "ne"
	CheckNameOneOf    = "oneof"
	CheckNameUnique   = "unique"

	CheckNameEmail = "email"
	CheckNameE164  = "e164"
	CheckNameUUID  = "uuid"

	CheckNameURI = "uri"
	CheckNameURL = "url"

	CheckNameHostname        = "hostname"         // RFC 952
	CheckNameHostnameRFC1123 = "hostname_rfc1123" // RFC 1123, DNS name
	CheckNameHostnamePort    = "hostname_port"    // [RFC 1123]:<port>
	CheckNameFQDN            = "fqdn"             // RFC 1123, but must contain a non-numerical TLD

	CheckNameRegexp = "regexp"

	CheckNameBase64       = "base64"
	CheckNameBase64URL    = "base64url"
	CheckNameBase64RawURL = "base64rawurl"

	CheckNameAlpha        = "alpha"
	CheckNameAlphanumeric = "alphanum"
	CheckNameNumeric      = "numeric"
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

func ErrRequired(name string) Error {
	return Error{CheckName: CheckNameRequired, ValueName: name}
}

func (err Error) Is(target error) bool {
	var validationErr Error

	if !errors.As(target, &validationErr) {
		return false
	}

	validationErr.Value = nil
	err.Value = nil

	return err == validationErr
}

func (err Error) As(target any) bool {
	if yavErr, ok := target.(*Error); ok && yavErr != nil {
		*yavErr = err
		return true
	}

	return false
}

func (err Error) Error() string {
	if err.CheckName != "" && err.ValueName != "" {
		if err.Parameter != "" {
			return fmt.Sprintf("'%s' failed the '%s=%q' check", err.ValueName, err.CheckName, err.Parameter)
		}

		return fmt.Sprintf("'%s' failed the '%s' check", err.ValueName, err.CheckName)
	}

	if err.CheckName != "" {
		return fmt.Sprintf("validation failed: %s", err.CheckName)
	}

	return "unknown validation error"
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
	case bool:
		return strconv.FormatBool(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}

	return ""
}

type Errors struct {
	Unknown    []error
	Validation []Error
}

func (errs Errors) Is(target error) bool {
	if yavErrs, ok := target.(Errors); ok && errs.IsZero() && yavErrs.IsZero() {
		return true
	}

	for _, err := range errs.Unknown {
		if errors.Is(err, target) {
			return true
		}
	}

	for _, yavErr := range errs.Validation {
		if yavErr.Is(target) {
			return true
		}
	}

	return false
}

func (errs Errors) As(target any) bool {
	if yavErrs, ok := target.(*Errors); ok && yavErrs != nil {
		*yavErrs = errs
		return true
	}

	for _, err := range errs.Unknown {
		if errors.As(err, target) {
			return true
		}
	}

	for _, yavErr := range errs.Validation {
		if yavErr.As(target) {
			return true
		}
	}

	return false
}

func (errs Errors) IsZero() bool {
	return len(errs.Unknown)+len(errs.Validation) == 0
}

func (errs Errors) Error() string {
	switch len(errs.Unknown) + len(errs.Validation) {
	case 0:
		return ""
	case 1:
		if len(errs.Unknown) > 0 {
			return errs.Unknown[0].Error()
		}

		return errs.Validation[0].Error()
	default:
		return errs.error()
	}
}

func (errs Errors) error() string {
	s := bytes.NewBuffer(make([]byte, 0, (len(errs.Unknown)+len(errs.Validation))*50))

	for _, err := range errs.Unknown {
		if s.Len() > 0 {
			s.WriteString("; ")
		}

		s.WriteString(err.Error())
	}

	for _, err := range errs.Validation {
		if s.Len() > 0 {
			s.WriteString("; ")
		}

		s.WriteString(err.Error())
	}

	return s.String()
}

// AsError returns non-zero Errors or nil.
func (errs Errors) AsError() error {
	if errs.IsZero() {
		return nil
	}

	return errs
}

func (errs *Errors) Append(err error) {
	if err == nil {
		return
	}

	switch typedErr := err.(type) {
	case Errors:
		if len(errs.Unknown) > 0 {
			errs.Unknown = append(errs.Unknown, typedErr.Unknown...)
		} else {
			errs.Unknown = typedErr.Unknown
		}

		if len(errs.Validation) > 0 {
			errs.Validation = append(errs.Validation, typedErr.Validation...)
		} else {
			errs.Validation = typedErr.Validation
		}
	case Error:
		errs.Validation = append(errs.Validation, typedErr)
	default:
		errs.Unknown = append(errs.Unknown, typedErr)
	}
}
