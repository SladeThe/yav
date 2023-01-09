package yav

import (
	"log"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
	"go.uber.org/multierr"

	"github.com/SladeThe/yav/common"
)

var (
	playgroundValuelessCheckNames = map[string]struct{}{
		CheckNameRequired:           {},
		CheckNameRequiredWithAny:    {},
		CheckNameRequiredWithoutAny: {},
		CheckNameRequiredWithAll:    {},
		CheckNameRequiredWithoutAll: {},
	}

	playgroundParameterlessCheckNames = map[string]struct{}{
		CheckNameRequired: {},

		CheckNameUnique: {},

		CheckNameEmail: {},
		CheckNameE164:  {},
		CheckNameUUID:  {},

		CheckNameLowercase: {},
		CheckNameUppercase: {},

		CheckNameContainsAlpha:      {},
		CheckNameContainsLowerAlpha: {},
		CheckNameContainsUpperAlpha: {},
		CheckNameContainsDigit:      {},

		CheckNameStartsWithAlpha:      {},
		CheckNameStartsWithLowerAlpha: {},
		CheckNameStartsWithUpperAlpha: {},
		CheckNameStartsWithDigit:      {},

		CheckNameEndsWithAlpha:      {},
		CheckNameEndsWithLowerAlpha: {},
		CheckNameEndsWithUpperAlpha: {},
		CheckNameEndsWithDigit:      {},

		CheckNameText:  {},
		CheckNameTitle: {},
	}
)

type Playground struct {
	Validator *validator.Validate
}

func NewPlayground() Playground {
	p := Playground{Validator: validator.New()}

	// When validation fails, use json tag name instead of Go struct field name.
	p.useJSONFieldName()

	// Register additional validation tags.
	p.mustRegisterContains()
	p.mustRegisterExcludes()
	p.mustRegisterStartsWith()
	p.mustRegisterEndsWith()
	p.mustRegisterText()
	p.mustRegisterTitle()

	return p
}

func (p Playground) Validate(s any) error {
	err := p.Validator.Struct(s)
	if err == nil {
		return nil
	}

	if fieldErrs, ok := err.(validator.ValidationErrors); ok {
		var combinedErr error

		for _, fieldErr := range fieldErrs {
			tag := fieldErr.Tag()

			parameter := fieldErr.Param()
			if _, omitParameter := playgroundParameterlessCheckNames[tag]; omitParameter {
				parameter = ""
			}

			value := fieldErr.Value()
			if _, omitValue := playgroundValuelessCheckNames[tag]; omitValue {
				value = nil
			}

			multierr.AppendInto(&combinedErr, Error{
				CheckName: tag,
				Parameter: parameter,
				ValueName: fieldErr.Field(),
				Value:     value,
			})
		}

		return combinedErr
	}

	return Error{}
}

func (p Playground) useJSONFieldName() {
	p.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

func (p Playground) mustRegisterContains() {
	p.Validator.RegisterAlias(CheckNameContainsAlpha, "containsany="+common.CharactersAlpha)
	p.Validator.RegisterAlias(CheckNameContainsLowerAlpha, "containsany="+common.CharactersLowerAlpha)
	p.Validator.RegisterAlias(CheckNameContainsUpperAlpha, "containsany="+common.CharactersUpperAlpha)
	p.Validator.RegisterAlias(CheckNameContainsDigit, "containsany="+common.CharactersDigit)

	fn := func(fl validator.FieldLevel) bool {
		s, ok := p.fieldAsString(fl)
		if !ok || s == "" {
			return false
		}

		return strings.ContainsAny(s, common.CharactersSpecial)
	}

	if err := p.Validator.RegisterValidation(CheckNameContainsSpecialCharacter, fn); err != nil {
		log.Fatal(err)
	}
}

func (p Playground) mustRegisterExcludes() {
	fn := func(fl validator.FieldLevel) bool {
		s, ok := p.fieldAsString(fl)
		if !ok {
			return false
		}

		for _, r := range s {
			if unicode.IsSpace(r) {
				return false
			}
		}

		return true
	}

	if err := p.Validator.RegisterValidation(CheckNameExcludesWhitespace, fn); err != nil {
		log.Fatal(err)
	}
}

func (p Playground) mustRegisterStartsWith() {
	checks := []struct {
		name string
		do   func(r rune) bool
	}{{
		name: CheckNameStartsWithAlpha,
		do:   common.IsAlpha,
	}, {
		name: CheckNameStartsWithLowerAlpha,
		do:   common.IsLowerAlpha,
	}, {
		name: CheckNameStartsWithUpperAlpha,
		do:   common.IsUpperAlpha,
	}, {
		name: CheckNameStartsWithDigit,
		do:   common.IsDigit,
	}, {
		name: CheckNameStartsWithSpecialCharacter,
		do:   common.IsSpecialCharacter,
	}}

	for _, check := range checks {
		do := check.do

		fn := func(fl validator.FieldLevel) bool {
			s, ok := p.fieldAsString(fl)
			if !ok || s == "" {
				return false
			}

			r, _ := utf8.DecodeRuneInString(s)
			return do(r)
		}

		if err := p.Validator.RegisterValidation(check.name, fn); err != nil {
			log.Fatal(err)
		}
	}
}

func (p Playground) mustRegisterEndsWith() {
	checks := []struct {
		name string
		do   func(r rune) bool
	}{{
		name: CheckNameEndsWithAlpha,
		do:   common.IsAlpha,
	}, {
		name: CheckNameEndsWithLowerAlpha,
		do:   common.IsLowerAlpha,
	}, {
		name: CheckNameEndsWithUpperAlpha,
		do:   common.IsUpperAlpha,
	}, {
		name: CheckNameEndsWithDigit,
		do:   common.IsDigit,
	}, {
		name: CheckNameEndsWithSpecialCharacter,
		do:   common.IsSpecialCharacter,
	}}

	for _, check := range checks {
		do := check.do

		fn := func(fl validator.FieldLevel) bool {
			s, ok := p.fieldAsString(fl)
			if !ok || s == "" {
				return false
			}

			r, _ := utf8.DecodeLastRuneInString(s)
			return do(r)
		}

		if err := p.Validator.RegisterValidation(check.name, fn); err != nil {
			log.Fatal(err)
		}
	}
}

// mustRegisterText registers common.IsText as a string value validator.
func (p Playground) mustRegisterText() {
	fn := func(fl validator.FieldLevel) bool {
		s, ok := p.fieldAsString(fl)
		return ok && common.IsText(s)
	}

	if err := p.Validator.RegisterValidation(CheckNameText, fn); err != nil {
		log.Fatal(err)
	}
}

// mustRegisterTitle registers common.IsTitle as a string value validator.
func (p Playground) mustRegisterTitle() {
	fn := func(fl validator.FieldLevel) bool {
		s, ok := p.fieldAsString(fl)
		return ok && common.IsTitle(s)
	}

	if err := p.Validator.RegisterValidation(CheckNameTitle, fn); err != nil {
		log.Fatal(err)
	}
}

func (p Playground) RegisterRegexp(name, pattern string) error {
	rgx, errCompile := regexp.Compile(pattern)
	if errCompile != nil {
		return errCompile
	}

	fn := func(fl validator.FieldLevel) bool {
		s, ok := p.fieldAsString(fl)
		return ok && rgx.MatchString(s)
	}

	return p.Validator.RegisterValidation(name, fn)
}

func (p Playground) MustRegisterRegexp(name, pattern string) {
	if err := p.RegisterRegexp(name, pattern); err != nil {
		log.Fatal(err)
	}
}

func (p Playground) fieldAsString(fl validator.FieldLevel) (s string, ok bool) {
	field := fl.Field()
	if field.Kind() != reflect.String {
		return "", false
	}

	return field.String(), true
}
